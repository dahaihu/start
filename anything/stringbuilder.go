package anything

import (
	"fmt"
	"strings"
)

type QueryArgs map[string]interface{}

func (q QueryArgs) Condition() (sql string, args []interface{}) {
	builder := strings.Builder{}
	cnt := 0
	for k, v := range q {
		if cnt == 0 {
			builder.WriteString("where ")
		} else {
			builder.WriteString(" and ")
		}
		cnt += 1

		builder.WriteString(fmt.Sprintf("`%s`=", k))
		args = append(args, v)
		switch v.(type) {
		case int64:
			builder.WriteString("?")
		case string:
			builder.WriteString("'?'")
		default:
			builder.WriteString("?")
		}
	}
	return builder.String(), args
}

func TestBuilder() {
	q := QueryArgs{"status": 1, "department_id": 10, "name": "name"}
	sql, args := q.Condition()
	fmt.Printf("sql is %v, args is %v\n", sql, args)
}
