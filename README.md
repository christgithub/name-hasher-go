# name-hasher
ETL pipeline hashing strings and indexing result in Elasticsearch

To make sure no container is running
```
docker stop $(docker ps -q)
```

To build the binary file
```
make build
```
To bring all containers up and start the program
```
make up
```

To verify the hashed entries in Elasticsearch, run:
```
docker exec -it elasticsearch /bin/sh
```

At the prompt, enter:
```
curl http://localhost:9200/sftp_index/_search?pretty=true
```

This should display a similar ouput, show the 10 first indexed hashed values and their source files
```
{
  "took" : 5,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 450,
      "relation" : "eq"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "sftp_index",
        "_id" : "CgX8LpgB8EyoSGOhwRko",
        "_score" : 1.0,
        "_source" : {
          "content" : "9e2d43f55514924202ce4c6d3961149f5c4e3e726c583bdcd7cab2c77fd8f5c5",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "DgX8LpgB8EyoSGOhwRlU",
        "_score" : 1.0,
        "_source" : {
          "content" : "a569dd2898808de896020fb467649fc5620992f6ed913e877302781298f74e46",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "DwX8LpgB8EyoSGOhwRlU",
        "_score" : 1.0,
        "_source" : {
          "content" : "fc9651de5bc42d3127d7a44806c34d87b27efc2d62b73c7a32d8d7a39ee0d624",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "EQX8LpgB8EyoSGOhwRlW",
        "_score" : 1.0,
        "_source" : {
          "content" : "8f67993675868c162fb62c38b1771ead36bfdae544fa540cf13f22a75bdeaa75",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "EwX8LpgB8EyoSGOhwRle",
        "_score" : 1.0,
        "_source" : {
          "content" : "70e046773bd6bded0f68bcd32df7cefebdadaba3b4f441f0b8bf28fbccf14afa",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "FAX8LpgB8EyoSGOhwRle",
        "_score" : 1.0,
        "_source" : {
          "content" : "3447bc83102ab46498d06b2b0301fe8bf005ddec72d2091fb64282bf66875fcd",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "FQX8LpgB8EyoSGOhwRlf",
        "_score" : 1.0,
        "_source" : {
          "content" : "b33912f676f7c0effa1e29d004ec1234df42c82decb684d09a5ac046ded1ca86",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "FgX8LpgB8EyoSGOhwRlj",
        "_score" : 1.0,
        "_source" : {
          "content" : "a423429dce82f620860b0eb089b3659f3835773c43d2c1268160d586013293a6",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "FwX8LpgB8EyoSGOhwRll",
        "_score" : 1.0,
        "_source" : {
          "content" : "1a36f0e66058de4f44418a8b7cac9e2ef67d0a49efb4050ff06b90db7c4e7d99",
          "filename" : "downloaded/input2.txt"
        }
      },
      {
        "_index" : "sftp_index",
        "_id" : "GAX8LpgB8EyoSGOhwRlm",
        "_score" : 1.0,
        "_source" : {
          "content" : "e29160b8d48afe37149d1063cd3ff2d7a51d37578b061a2c2af81d774df76bd8",
          "filename" : "downloaded/input2.txt"
        }
      }
    ]
  }
}
```

