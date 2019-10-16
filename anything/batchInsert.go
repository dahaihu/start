package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-08-06 14:04
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func BatchInsert(data []map[string]interface{}) {
	sqlStr := "INSERT INTO test(n1, n2, n3) VALUES "
	vals := make([]interface{}, 0)

	for _, row := range data {
		sqlStr += "(?, ?, ?),"
		vals = append(vals, row["v1"], row["v2"], row["v3"])
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	fmt.Println(sqlStr)
	fmt.Println(vals)
}


//func (d MysqlDao) BatchInsert(ctx context.Context, dao Mysql, ptrOfEntitys interface{}, ) error {
//	items := ptrOfEntitys.([]interface{})
//	builder := strings.Builder{}
//	builder.WriteString(fmt.Sprintf("INSERT INTO `%s` ", dao.Table()))
//	builder.WriteString(" (")
//	// 实体
//	elem := reflect.ValueOf(items[0]).Elem()
//	// 类型
//	resType := reflect.TypeOf(elem.Interface())
//
//	cnt := 0
//	for i := 0; i < elem.NumField(); i++ {
//		field := resType.Field(i)
//		mysqlCol := field.Tag.Get("json")
//
//		if _, ok := InsertWhiteFiled[mysqlCol]; ok {
//			continue
//		}
//		if cnt == 0 {
//			builder.WriteString(fmt.Sprintf("`%s`", mysqlCol))
//		} else {
//			builder.WriteString(fmt.Sprintf(",`%s`", mysqlCol))
//		}
//		cnt += 1
//	}
//	builder.WriteString(") VALUES ")
//	for i := 0; i < len(items); i++ {
//		for idx := 0; idx < elem.NumField(); idx++ {
//			if idx == 0 {
//				builder.WriteString("(?")
//			} else if idx == elem.NumField()-1 {
//				builder.WriteString("),")
//			} else {
//				builder.WriteString(",?")
//			}
//		}
//	}
//
//	args := make([]interface{}, 0)
//	for count := 0; count < len(items); count++ {
//		elem = reflect.ValueOf(items[count]).Elem()
//		for i := 0; i < elem.NumField(); i++ {
//			field := resType.Field(i)
//			mysqlCol := field.Tag.Get("json")
//			if _, ok := InsertWhiteFiled[mysqlCol]; ok {
//				continue
//			}
//			args = append(args, elem.Field(i).Interface())
//		}
//	}
//	fmt.Println(fmt.Sprintf("sql is %v", builder.String()[:len(builder.String())-1]))
//	return nil
//	//r, err := dao.Pool().Master().Exec(ctx, builder.String()[:len(builder.String())-1], args...)
//	//if err != nil {
//	//	return err
//	//}
//	//return err
//}
