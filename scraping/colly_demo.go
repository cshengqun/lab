package scraping

import (
	"encoding/json"
	"fmt"
	"github.com/cshengqun/lab/scraping/value_object"
	"github.com/gocolly/colly"
	"strings"
)


func CollyGetHtml() {
	c := colly.NewCollector()
	c.OnHTML(".info", func(e *colly.HTMLElement){
		// e.Request.Visit(e.Attr("href"))
		test := e.ChildText(".rating_num")
		fmt.Printf("test:%s\n", test)
		titla := e.ChildText(".hd a span:first-child")
		fmt.Printf("titla:%s\n", titla)
		return
		name := e.DOM.Find(".hd a").Text()
		rate := e.DOM.Find(".rating_num").Text()
		texts := strings.Split(name, "/")
		if len(texts) > 0 {
			fmt.Printf("%s %s\n", strings.TrimSpace(texts[0]), rate)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("response:", string(response.Body))
	})
	for start:=0; start<=250; start= start + 25 {
		url := fmt.Sprintf("https://movie.douban.com/top250?start=%d&filter=", start)
		c.Visit(url)
	}

}

func CollyPostBE(){
	var question []string
	res := &value_object.BdQuestion{}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request){
		fmt.Printf("%s %s\n", r.Method, r.URL)
		r.Headers.Set("X-Kunlun-Language-Code", "2052")
		r.Headers.Set("x-kunlun-fe-version", "8b7980788214__bc82c6ea-994c-405f-9dfd-f22f0d427b5c")
		r.Headers.Set("Cookie", "X-Kunlun-SessionId=G%3Afea1dc17e57c4e78bea4.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2YWwiOnsidGVuYW50X2lkIjozOTAsInVzZXJfaWQiOjE3MDM0NDA1OTcyMjc1NjAsInRlbmFudF9kb21haW5fbmFtZSI6IiIsInNlc3Npb25fdmVyc2lvbiI6InYyMDIwLTA1LTE5Iiwid3NfdG9rZW4iOiJXOmE0MzNkMTZlNTAwZDRmNmJhYmU1IiwibG9naW5fdG9rZW4iOiIxOGM3ODM2N2Y5NGY5MTRjaWdzZjk1ZGExNWU1NWI1MTQ2YTF5SmE2YWVhNjViNTYzYTUzOWU5T0dkYjllOTk4Y2YzNzU5ZGQxYU85In0sImV4cCI6MTY3NTE0MTI3OX0.3xx7MMtgq3lMAkjTTCN_sPE_CfUmz5bhlvkTdVW4Guw; MONITOR_WEB_ID=436f5ccbd6b7f1e84c398092d51819fd; _tea_utm_cache_1229=undefined")
		r.Headers.Set("X-Kunlun-Token", "18c78367f94f914cigsf95da15e55b5146a1yJa6aea65b563a539e9OGdb9e998cf3759dd1aO9")
	})
	c.OnResponse(func(response *colly.Response) {
		// fmt.Printf("res: %s\n", string(response.Body))
		err := json.Unmarshal(response.Body, res)
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}
		//fmt.Printf("%v\n", res)
		//fmt.Printf("plus:%+v\n", res)
		for _, item := range res.Data.Items {
			question = append(question, item.Basic.Name[0].Text)
		}

	})
	c.OnScraped(func(response *colly.Response) {
		/*
		for i, q := range question {
			fmt.Printf("%d.%s\n", i, q)
		}
		 */
		//fmt.Printf("on scraped\n")
	})
	for i:=1;i<35;i=i+1{
		cData := fmt.Sprintf(data, i)
		c.PostRaw("https://apaas.feishuapp.cn/ae/api/v2/layout/namespaces/Interview_Question_Community__c/table/runtime/data", []byte(cData))

	}
	for i, q := range question {
		fmt.Printf("%d.%s\n", i, q)
	}
}

var data string = `{
  "size": 10,
  "page": %d,
  "actions": [
    {
      "objectId": 1710959217315853,
      "actionId": 1710968160984125,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    },
    {
      "objectId": 1710959217315853,
      "actionId": 1710968343254056,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    },
    {
      "objectId": 1710959217315853,
      "actionId": 1710972167226430,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    },
    {
      "objectId": 1710959217315853,
      "actionId": 1711668869351543,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    },
    {
      "objectId": 1710959217315853,
      "actionId": 1711673997531160,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    },
    {
      "objectId": 1710959217315853,
      "actionId": 1717128548154375,
      "field": {
        "fieldApiName": "_id",
        "fieldLabel": [],
        "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "objectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ],
        "targetObjectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2",
        "targetObjectLabel": [
          {
            "language_code": 1033,
            "text": "Question"
          },
          {
            "language_code": 2052,
            "text": "题目"
          }
        ]
      }
    }
  ],
  "fields": [
    "_name",
    "Interview_Question_Community__c__lookup_aadbihau3kycq",
    "Interview_Question_Community__c__lookup_aadbig7sp6gb4",
    "Interview_Question_Community__c__lookup_aadbkporu4kbq",
    "Interview_Question_Community__c__number_aadbihau3kydq",
    "Interview_Question_Community__c__number_aadbihakeeed4",
    "Interview_Question_Community__c__text_aadbiufv5xgaq",
    "_id"
  ],
  "objectId": 1710959217315853,
  "authorize": {
    "appId": 1710957518900238,
    "appPageName": "9ppivpl8cubeoa_r31cb21"
  },
  "pageVariableMap": {
    "var_hqm1scf": {
      "type": "constant",
      "value": false
    },
    "var_hdkvk0e": {
      "type": "constant",
      "recordIds": [],
      "objectId": 1710959217315853
    },
    "var_4gdf2jc": {
      "type": "constant",
      "recordIds": [
        1713100974333959
      ],
      "objectId": 1710957331096663
    },
    "var_l00tas5": {
      "type": "constant",
      "value": "Interview_Question_Community__c__option_1631701284677_gh5emow"
    },
    "var_d11u86n": {
      "type": "constant",
      "value": "Interview_Question_Community__c__option_1631701284677_gh5emow"
    }
  },
  "sort": [
    {
      "direction": "desc",
      "field": "Interview_Question_Community__c__number_aadbjpx74judq",
      "isFormula": false,
      "type": "number"
    },
    {
      "direction": "desc",
      "field": "Interview_Question_Community__c__number_aadbjpyaqgsfq",
      "isFormula": false,
      "type": "number"
    },
    {
      "direction": "asc",
      "field": "Interview_Question_Community__c__number_aadbjpybqkcbq",
      "isFormula": false,
      "type": "number"
    }
  ],
  "describeFilter": {
    "filter": {
      "api_name": "",
      "biz_type": "",
      "conditions": [
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__option_aadbihakeega4",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "enum"
            },
            "type": "metadataVariable"
          },
          "operator": "notEquals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "data": "Interview_Question_Community__c__option_1631700928229_228llva"
            },
            "type": "constant"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__option_aadbihb2pw4dq",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "enum"
            },
            "type": "metadataVariable"
          },
          "operator": "equals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "data": "Interview_Question_Community__c__option_1631701381690_1ojabpa"
            },
            "type": "constant"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__boolean_aadbixhtkjuh4",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "boolean"
            },
            "type": "metadataVariable"
          },
          "operator": "equals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "variableApiName": "var_hqm1scf",
              "variableType": "boolean"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__boolean_aadbixhtkjuh4",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "boolean"
            },
            "type": "metadataVariable"
          },
          "operator": "notEquals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "variableApiName": "var_hqm1scf",
              "variableType": "boolean"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "_id",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          },
          "operator": "isAnyOf",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "_id",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "var_hdkvk0e",
              "variableType": "lookup__multiple"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [
                    "all"
                  ],
                  "fieldApiName": "Interview_Question_Community__c__lookup_aadbiwjje4sd2",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          },
          "operator": "hasAnyOf",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "_id",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig2q7nugo"
                }
              ],
              "variableApiName": "var_4gdf2jc",
              "variableType": "lookup__multiple"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__option_aadbihb2pdeao",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "enum"
            },
            "type": "metadataVariable"
          },
          "operator": "equals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "variableApiName": "var_l00tas5",
              "variableType": "enum"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__option_aadbihb2pdeao",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "enum"
            },
            "type": "metadataVariable"
          },
          "operator": "equals",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "variableApiName": "var_d11u86n",
              "variableType": "enum"
            },
            "type": "customizedVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__lookup_aadbnvuknhgm2",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                },
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__lookup_aadbntslummc2",
                  "objectApiName": "Interview_Question_Community__c__object_aadboac4jjqco"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          },
          "operator": "hasAllOf",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "_id",
                  "objectApiName": "_user"
                }
              ],
              "variableApiName": "current_user",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          }
        },
        {
          "index": 0,
          "left": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__lookup_aadbnvuknhgm2",
                  "objectApiName": "Interview_Question_Community__c__object_aadbig54l6gd2"
                },
                {
                  "extendLogicTags": [],
                  "fieldApiName": "Interview_Question_Community__c__lookup_aadboakphjaf2",
                  "objectApiName": "Interview_Question_Community__c__object_aadboac4jjqco"
                }
              ],
              "variableApiName": "current_object",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          },
          "operator": "hasAllOf",
          "right": {
            "displayName": null,
            "fieldsDisplayNames": null,
            "settings": {
              "fieldPath": [
                {
                  "extendLogicTags": [],
                  "fieldApiName": "_id",
                  "objectApiName": "_user"
                }
              ],
              "variableApiName": "current_user",
              "variableType": "lookup"
            },
            "type": "metadataVariable"
          }
        }
      ],
      "logic": "(1 and (2 or (9 or 10)) and (3 or (4 and 5)) and 6 and (7 or 8))",
      "use_type": "",
      "version": "v2"
    }
  },
  "fromComponent": "list"
}`
