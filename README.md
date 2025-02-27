## go-backend-vetautet

## How to run 

Trong folder có thư mục `environment` ở đây chứa tất cả các Stacks công nghệ bạn phải nằm thông qua như
`Kafka`, `redis`, `master/slave`, `sentinel`, `RateLimiter`, `CircuitBreaker`, `MySQL`...

Chú ý ở đây có hai file `docker-compose` nếu các bạn máy có cấu hình trung bình, vui lòng 

```bash
~ docker-compose -f environment/docker-compose-dev.yml up
```

Thì docker sẽ run những container cần thiết cho dự án.

Ngược lại nếu máy của bạn khác mạnh thì có thể run bản full của docker.

```bash
~ docker-compose -f environment/docker-compose-pro.yml up
```

Ngoài ra hãy chuyên nghiệp khi sử dụng `makefile`.

```bash
docker_up:
	docker-compose -f environment/docker-compose-dev.yml up
docker_down:
	docker-compose -f environment/docker-compose-dev.yml down
```

## How to test

```bash
~ curl http://localhost:8002/v1/2024/ticket/item/4
```

Output:

```go 
{
    "code":20001,
    "message":"success",
    "data":
        {
            "ticket_Id":1,
            "ticket_Name":
            "Vé Sự Kiện 12/12 - Hạng Phổ Thông",
            "stock_Available":1000,
            "stock_Initial":1000
        }
    }
```
### Mentor: anonystick