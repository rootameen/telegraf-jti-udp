#
# Copyright (c) 2017 Juniper Networks, Inc. All rights reserved.
#

##############################
## Supporting functions     ##
##############################

package main

import(

	"fmt"
	
)


func parse_check(data map[string]interface{}) string {
        for key,val := range data{
                switch v := val.(type) {
                case int:
                        fmt.Printf("Value = %s, Type = %s , key=%s, v=%s\n", val, reflect.TypeOf(val), key, v)
                        return "int"
                case float64:
                        fmt.Printf("Value = %s, Type = %s \n", val, reflect.TypeOf(val))
                        return "float"
                case string:
                        fmt.Printf("Value = %s, Type = %s \n", val, reflect.TypeOf(val))
                        return "str"
                case []interface{}:
                        for i,e := range(v){
                                fmt.Printf("\n@@@@@@@@%s: %s %s\n", i,reflect.TypeOf(e), e)
                                //parse_check(e.(map[string]interface{}))
                        }
                        fmt.Printf("Entered array of interfaces!! \n")
                default:
                  // i isn't one of the types above
                        data_inner, ok := val.(map[string]interface{})
                        if !ok {
                                 panic("inner map is not a map!!!")
                        }
                        parse_check(data_inner)

                }

        }
        return "end"
}




func parse_map(data map[string]interface{}, master_key string) []interface{} {
	var leaf_data map[string]interface{}
	var arr_data []interface{}
	var arr_key []string
	var fin_data []interface{}
	
	for key,val := range data{
		if master_key == ""{
			new_master_key = key	
		} else {
			new_master_key = master_key + "." + key
		}
		
		valType := reflect.ValueOf(val).Kind()
		if valType == reflect.map{
			map_data := parse_map(val, new_master_key)
			fmt.Printf("Map data = %s \n", map_data)
			arr_data = append(arr_data, map_data)
			arr_key = append(arr_key, new_master_key)
		} else if valType == reflect.Slice{
			arr_data = append(arr_data, parse_array(val, new_master_key))
			arr_key = append(arr_key, new_master_key)

		} else { leaf_data[new_master_key] = val}

//		switch s := val.(type) {
//		case int:
//			leaf_data[new_master_kay] = val
//		case float64:
//			leaf_data[new_master_key] = val
//		case string:
//			leaf_data[new_master_key] = val
//		case []interface{}:
//			arr_data = append(arr_data, parse_array(value, new_master_key))	
//			arr_key = append(arr_key, new_master_key)
//		case map[string]interface{}:
//			map_data := parse_map(val, new_master_key)
//			if reflect.ValueOf(map_data).Kind() == reflect.Map {arr_data = append(arr_data, [map_data])}
//		
//		default:
//			fmt.Println("Cannot understand value type")
//			fmt.Print("Type = %s", s)
//
//		}
		
	}

	if len(leaf_data) == 0 {
		for i,key := range len(arr_key){
			for _, data_aa := arr_data[i]{
				leaf_tmp := leaf_data
				if data_aa != nil {
					if reflect.ValueOf(data_aa).Kind() == reflect.map{
						for key_aa, val_aa := range data_aa{
							leaf_tmp[key_aa] = value_aa
						}
						fin_data = append(fin_data, leaf_tmp)
					} else { 
						 for _,data_ha := range	data_aa{
							leaf_tmp := leaf_data
							for key_aa,value_aa := range data_ha{
								leaf_tmp[key_aa] = value_aa
							}
							fin_data = append(fin_data, leaf_tmp)
						 }
						}
				}
			}

		}
	
	} else fin_data = arr_data

	if len(fin_data == 0 && len(leaf_data)!= 0){
		
		fin_data = append(fin_data, leaf_data)
	}

	return fin_data
}

func parse_array(data []interface{}, master_key string){
	var arr_data []interface{}
	for val := range data{
		valType := reflect.ValueOf(val).Kind()
		if valType == reflect.map{
			arr_data = append(arr_data, parse_map(val, master_key))
		} else {fmt.Println("Error!!!! Leaf elements in array are not coded. Please open a issue.")}	i
	}
	
	return arr_data
}


