# Домашнее задание к занятию "7.6. Написание собственных провайдеров для Terraform."

## Задача 1. 

1. Найдите, где перечислены все доступные `resource` и `data_source`, приложите ссылку на эти строки в коде на 
гитхабе.   

>Все доступные `resource` перечислены в файле [provider.go, строки с 914](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/provider/provider.go#L914)
>
>Все доступные `data_source` перечислены в файле [provider.go, строки с 425](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/provider/provider.go#L425)
 
 
2. Для создания очереди сообщений SQS используется ресурс `aws_sqs_queue` у которого есть параметр `name`.

* С каким другим параметром конфликтует `name`? Приложите строчку кода, в которой это указано.
 
```hcl
		"name": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ForceNew:      true,
			ConflictsWith: []string{"name_prefix"},
		},
```
> [ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L87)

* Какая максимальная длина имени? 

> 80 символов
> 
>[ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L427)
  
* Какому регулярному выражению должно подчиняться имя? 
    
```hcl
if fifoQueue {
                        re = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,75}\.fifo$`)
                } else {
                        re = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,80}$`)
                }
```
>[ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L425)
> 
>[ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L427)
