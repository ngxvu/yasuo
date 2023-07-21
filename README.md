
# Meradia

![flow.png](flow.png)

Link: https://merakilab-space.notion.site/Update-image-flow-c4d8b7ac4f4842049d7162621943db45?pvs=4

Để upload ảnh lên S3 storage bằng ReactJS và Golang, bạn có thể làm theo các bước sau:

### **Bước 1: Gọi Backend để lấy Presigned URL**

- Trong frontend (ReactJS), gửi một yêu cầu HTTP tới backend (Golang) để lấy presigned URL. Đảm bảo rằng backend đã cấu hình và có quyền truy cập vào Amazon S3.
- Backend sẽ nhận yêu cầu, sử dụng credential của nó để kết nối tới Amazon S3 và tạo một presigned URL. Presigned URL này sẽ được tạo với quyền truy cập cho việc upload (ACL) và có thời hạn.
- Backend sẽ trả về presigned URL cho frontend.

### **Bước 2: Tải ảnh lên S3 bằng Presigned URL**

- Trong frontend, sử dụng presigned URL đã nhận được để tải ảnh lên S3. Bạn có thể sử dụng các thư viện như Axios hoặc Fetch để thực hiện yêu cầu HTTP POST tới presigned URL với ảnh dữ liệu đính kèm.
- Khi yêu cầu này được gửi đi, ảnh sẽ được truyền từ frontend trực tiếp tới S3 thông qua presigned URL.

### **Bước 3: Xác nhận và lưu thông tin vào Backend**

- Sau khi quá trình upload hoàn thành, frontend có thể gửi một yêu cầu khác tới backend để thông báo rằng quá trình upload đã hoàn tất. Yêu cầu này có thể chứa thông tin về ảnh đã upload hoặc bất kỳ thông tin khác liên quan.
- Backend có thể tiếp nhận yêu cầu này, xác nhận rằng quá trình upload đã hoàn thành và lưu thông tin về ảnh vào cơ sở dữ liệu hoặc bất kỳ hệ thống nào khác.
- Sau khi lưu thông tin thành công, backend có thể trả về một liên kết (link) tới ảnh đã lưu, để frontend có thể truy cập và hiển thị ảnh.

Đây là một quy trình cơ bản để upload ảnh lên S3 storage sử dụng ReactJS và Golang. Tuy nhiên, cần lưu ý rằng cài đặt chi tiết của từng bước có thể thay đổi tùy thuộc vào yêu cầu và cấu hình của bạn.


---
# APIs:

### PreUpload:
```curl
curl --location 'localhost:8099/api/v1/media/pre-upload' \
--header 'Content-Type: application/json' \
--data '{
    "media_type": "png"
}'
```

### Upload: BE test only
```curl
curl --location 'localhost:8099/api/v1/media/upload' \
--form 'upload_url="http://noormatch.ap-south-1.linodeobjects.com/noormatch/dde11abc-6257-44cd-be02-ac03a64359cc/image/159f15ba-004a-4bd0-86b9-fccb60959e5b.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=R53GYSC97373T37GXQZN%2F20230721%2Fap-south-1%2Fs3%2Faws4_request&X-Amz-Date=20230721T085505Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host%3Bx-amz-acl&x-id=PutObject&X-Amz-Signature=71ca596f7f6679e5f59031f4ced3401ad3b9eb94d25400854f408b84e7acf8fe"' \
--form 'file=@"/C:/Users/AlienWare/Downloads/3 copy 1.png"'
```

### Pos upload:
* get [name] from PreUpload api response
```curl
curl --location 'localhost:8099/api/v1/media/pos-upload' \
--header 'Content-Type: application/json' \
--data '{
    "name": "159f15ba-004a-4bd0-86b9-fccb60959e5b.png",
    "media_type": "png",
    "type_to_crop": "avatar"
}'
```