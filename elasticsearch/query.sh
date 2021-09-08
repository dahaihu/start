# query all mappings
curl 'http://localhost:9200/_mappings'

# query all records
curl -X GET -H 'Content-Type: application/json' -d '{"query":{"match_all":{}}}' 'http://localhost:9200/group/_doc/_search?pretty=true'

# fuzzy search
curl -X GET -H 'Content-Type: application/json' -d '{"query":{"fuzzy":{"name":"zhangsan"}}}' 'http://localhost:9200/schools/children/_search?pretty=true'

# query by id
curl 'http://localhost:9200/schools/children/LBuoUncB9ImVji_oGhKS'

# insert record
curl -X POST -H 'Content-Type: application/json' -d '{"name":"wangwu","age":30}' 'http://localhost:9200/schools/children'

# search multi data
curl -X GET -H 'Content-Type: application/json' -d '{"query":{"terms": {"age":[10, 20]}}}' 'http://localhost:9200/schools/children/_search?pretty=true'

# update data
curl -X POST -H 'Content-Type: application/json' -d '{"doc":{"name":"lisi"}}' 'http://localhost:9200/schools/children/LRu3UncB9ImVji_oDRIs/_update'

# add data to company/children
curl -X POST -H 'Content-Type: application/json' -d '{"name":"zhangsan","age":10,"father": {"names":{"old":"wangyi", "new":"wanger"},"ages":[1,2,3,4]}}' 'http://localhost:9200/company/children'

curl -X GET -H 'Content-Type: application/json' -d '{"query":{"match":{"father.names.old": "liyi"}}}' 'http://localhost:9200/company/children/_search?pretty=true'

# multi data
curl -X GET -H 'Content-Type: application/json' -d '{"query":{"terms": {"father.wori.old":["lisi", "wangyi"]}}}' 'http://localhost:9200/company/children/_search?pretty=true'

# range search

# group role
curl -X POST -H 'Content-Type: application/json' -d '{
  "group" : "xusong",
  "users" : [
    {
      "first" : "hu",
      "last" :  "san"
    },
    {
      "first" : "hu",
      "last" :  "si"
    }
  ]
}' 'http://localhost:9200/group/_doc'



# group role update
curl -X POST -H 'Content-Type: application/json' -d '{
    "script": {
        "lang": "painless",
        "source": "boolean exist=false; for(e in ctx._source.users){if (e.first == params.first) {e.last = params.last; exist=true; break;}} if (!exist) {ctx._source.users.add(params);}",
        "params": {
        	"first": "zhang",
        	"last": "si"
        }
    }
}' 'http://localhost:9200/group/_doc/pRNMVncBR3eyHNE8iMy3/_update'

curl -X DELETE 'http://localhost:9200/group'

# nested field create, update, delete
curl -X PUT -H 'Content-Type: application/json' -d '
{
    "mappings": {
         "properties" : {
             "user_role" : {
                 "type" : "nested",
                 "properties" : {
                     "user_id" : { "type" : "long" },
                     "role_id" : { "type" : "long" }
                 }
             }
         }
    }
}' 'http://localhost:9200/resource'

curl -X POST -H 'Content-Type: application/json' -d '{
    "name":"zhangsan",
    "user_role":[{
        "user_id":1,
        "role_id":2
    }]
}' 'http://localhost:9200/resource/_doc/1'
curl -XPOST -H 'Content-type: application/json' -d '{"script":{"lang": "painless","params":{"item":{"role_id":1,"user_id":10}},"source":"ctx._source.user_role.add(params.item)"}}' localhost:9200/resource/_doc/1/_update
curl -XPOST -H 'Content-type: application/json' -d '{"script":{"params":{"user_id":10},"source":"ctx._source.user_role.removeIf(item -\u003e item.user_id == params.user_id)"}}' localhost:9200/resource/_doc/1/_update
curl -XPOST -H 'Content-type: application/json' -d '{"script":{"params":{"role_id":100,"user_id":100},"source":"ctx._source.user_role.removeIf(item -\u003e item.user_id == params.user_id); ctx._source.user_role.add(params)"}}' localhost:9200/resource/_doc/1/_update

curl -X POST -H 'Content-Type: application/json' -d '{
  "tags" : "zhoujielun",
  "comments" : {"username": "zhangsan", "comment": "sb"}
}' 'http://localhost:9200/issues/comment'

# boke
curl -X PUT "localhost:9200/group" -H 'Content-Type: application/json' -d'
{
  "mappings": {
      "properties": {
        "users": {
          "type": "nested"
        }
      }
    }
}
'

# put index
curl -X PUT '
{
  "mappings": {
    "properties": {
      "users": {
        "type": "nested"
      }
    }
  }
}' 'http://localhost:9200/group'


# nested search
curl -X GET -d '{
  "query": {
    "bool": {
      "must": [
        {
          "nested": {
            "path": "users",
            "query": {
              "bool": {
                "should": [
                  {
                    "match": {
                      "users.first": "wang"
                    }
                  },
                  {
                    "match": {
                      "users.last": "2"
                    }
                  }
                ],
                "must": [
                  {
                    "match": {
                      "users.last": "si"
                    }
                  }
                ]
              }
            }
          }
        }
      ]
    }
  }
}' -H 'Content-Type: application/json' 'http://localhost:9200/group/_doc/_search'



# 父子关系索引的建立
curl -X PUT -d '{
  "mappings": {
    "properties": {
      "relation_type": {
        "type": "join",
        "eager_global_ordinals": true,
        "relations": {
          "project": "training",
          "training": "resource"
        }
      }
    }
  }
}' -H 'Content-Type: application/json' 'http://localhost:9200/project'


curl -X POST -d '
{
    "name":"resource4",
    "type": "resourceType4",
    "relation_type":{
        "name":"resource",
        "parent": "gS4pYncBzPzGau25rjx0"
    }
}
' -H 'Content-Type: application/json' 'http://localhost:9200/project/_doc?routing=fy4oYncBzPzGau25FjyU'


# 根据父节点，查询所有的子节点
curl -X GET -H 'Content-Type: application/json' -d '
{
  "query": {
    "parent_id": {
      "type": "training",
      "id": "fy4oYncBzPzGau25FjyU"
    }
  }
}
' 'http://localhost:9200/project/_doc/_search?pretty=true'


# 查询 title 包含 first 的父文档的所有子文档
curl -X GET -H 'Content-Type: application/json' -d '
{
  "query": {
    "has_parent": {
      "parent_type": "project",
      "query": {
        "match": {
          "name": "name"
=======
# parent child

curl -X PUT -H 'Content-Type: application/json' '
{
  "mappings": {
    "properties": {
      "my_id": {
        "type": "keyword"
      },
      "my_join_field": {
        "type": "join",
        "relations": {
          "question": "answer"
        }
      }
    }
  }
}


curl -X GET -H 'Content-Type: application/json' -d '
{
  "query": {
    "has_parent": {
      "parent_type": "training",
      "query": {
        "has_parent": {
          "parent_type": "project",
          "query": {
            "match": {
              "name": "name"
            }
          }
        }
      }
    }
  }
}
}' 'http://localhost:9200/my-index-000001'


curl -X PUT -H 'Content-Type: application/json' '
{
  "my_id": "1",
  "text": "This is a question",
  "my_join_field": {
    "name": "question"
  }
