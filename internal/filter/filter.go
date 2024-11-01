package filter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

const (
	Search   = 1 // Filter response with LIKE query "search={search_phrase}"
	Filter   = 2 // Filter response by column name values "filter={column_name}:{value}"
	Paginate = 4 // Paginate response with page and page_size
	OrderBy  = 8 // Order response by column name
	TagKey   = "filter"
)

var (
	paramNameRegexp = regexp.MustCompile(`(?m)param:(\w+).*`)
)

// queryParams represents the query parameters for filtering and pagination.
func orderBy(db *gorm.DB, params queryParams) *gorm.DB {
	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: params.OrderBy},
		Desc:   params.OrderDirection == "desc"},
	)
}

// filterByColumn filters the response by column name values.
func paginate(db *gorm.DB, params queryParams) *gorm.DB {
	if params.Page == 0 {
		params.Page = 1
	}

	switch {
	case params.PageSize > 100:
		params.PageSize = 100
	case params.PageSize <= 0:
		params.PageSize = 10
	}

	offset := (params.Page - 1) * params.PageSize
	return db.Offset(offset).Limit(params.PageSize)
}

// filterByColumn filters the response by column name values.
func searchField(columnName string, field reflect.StructField, phrase string) clause.Expression {
	filterTag := field.Tag.Get(TagKey)

	if strings.Contains(filterTag, "searchable") {
		return clause.Like{
			Column: clause.Expr{SQL: "LOWER(?)", Vars: []interface{}{clause.Column{Table: clause.CurrentTable, Name: columnName}}},
			Value:  "%" + strings.ToLower(phrase) + "%",
		}
	}
	return nil
}

// filterByColumn filters the response by column name values.
func filterField(columnName string, field reflect.StructField, phrase string) clause.Expression {
	var paramName string
	if !strings.Contains(field.Tag.Get(TagKey), "filterable") {
		return nil
	}
	paramMatch := paramNameRegexp.FindStringSubmatch(field.Tag.Get(TagKey))
	if len(paramMatch) == 2 {
		paramName = paramMatch[1]
	} else {
		paramName = columnName
	}

	// re, err := regexp.Compile(fmt.Sprintf(`(?m)%v([:<>!=]{1,2})(\w{1,}).*`, paramName))
	// for the current regex, the compound operators (such as >=) must come before the
	// single operators (such as <) or they will be incorrectly identified
	re, err := regexp.Compile(fmt.Sprintf(`(?m)%v(:|!=|>=|<=|>|<|~)([^,]*).*`, paramName))
	if err != nil {
		return nil
	}
	filterSubPhraseMatch := re.FindStringSubmatch(phrase)
	if len(filterSubPhraseMatch) == 3 {
		switch filterSubPhraseMatch[1] {
		case ">=":
			return clause.Gte{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		case "<=":
			return clause.Lte{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		case "!=":
			return clause.Neq{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		case ">":
			return clause.Gt{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		case "<":
			return clause.Lt{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		case "~":
			return clause.Like{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		default:
			return clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: columnName}, Value: filterSubPhraseMatch[2]}
		}
	}
	return nil
}

// expressionByField returns a gorm.DB with the filter applied.
func expressionByField(
	db *gorm.DB, phrases []string,
	operator func(string, reflect.StructField, string) clause.Expression,
	predicate func(...clause.Expression) clause.Expression,
) *gorm.DB {
	modelType := reflect.TypeOf(db.Statement.Model).Elem()
	numFields := modelType.NumField()
	modelSchema, err := schema.Parse(db.Statement.Model, &sync.Map{}, db.NamingStrategy)
	if err != nil {
		return db
	}
	var allExpressions []clause.Expression

	for _, phrase := range phrases {
		expressions := make([]clause.Expression, 0, numFields)
		for i := 0; i < numFields; i++ {
			field := modelType.Field(i)
			expression := operator(modelSchema.LookUpField(field.Name).DBName, field, phrase)
			if expression != nil {
				expressions = append(expressions, expression)
			}
		}
		if len(expressions) > 0 {
			allExpressions = append(allExpressions, predicate(expressions...))
		}
	}
	if len(allExpressions) == 1 {
		db = db.Where(allExpressions[0])
	} else if len(allExpressions) > 1 {
		db = db.Where(predicate(allExpressions...))
	}
	return db
}

// FilterByQuery filters the response by query parameters.
// Example:
//
//	db.Model(&UserModel).Scope(filter.FilterByQuery(ctx, filter.ALL)).Find(&users)
//
// Or if only pagination and order is needed:
//
//	db.Model(&UserModel).Scope(filter.FilterByQuery(ctx, filter.PAGINATION|filter.ORDER_BY)).Find(&users)
//
// And models should have appropriate`filter` tags:
//
//	type User struct {
//		gorm.Model
//		Username string `gorm:"uniqueIndex" filter:"param:login;searchable;filterable"`
//		// `param` defines custom column name for the query param
//		FullName string `filter:"searchable"`
//	}
func FilterByQuery(c *gin.Context, config int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var params queryParams
		err := c.BindQuery(&params)
		if err != nil {
			return db
		}

		model := db.Statement.Model
		modelType := reflect.TypeOf(model)
		if model != nil && modelType.Kind() == reflect.Ptr && modelType.Elem().Kind() == reflect.Struct {
			if config&Search > 0 && params.Search != "" {
				db = expressionByField(db, []string{params.Search}, searchField, clause.Or)
			}
			if config&Filter > 0 && len(params.Filter) > 0 {
				db = expressionByField(db, params.Filter, filterField, clause.And)
			}
		}

		if config&OrderBy > 0 {
			db = orderBy(db, params)
		}
		if config&Paginate > 0 {
			db = paginate(db, params)
		}
		return db
	}
}
