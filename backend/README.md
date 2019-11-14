**Screamdeal** Golang API Documentation

### Local setup

- Add .env file in root of /backend (ask Christoffer for the values)
- Run docker-compose up

### Endpoints

#### GET

**/themes**

#### POST

**/put-pdf**

return link to pdf in S3 (file storage)

takes parameters PDF file (somehow)

#### /add-theme-to-pdf

returns link to themed pdf in S3 (file storage)

takes parameters PDF link (from /put-pdf) and link to theme image (in S3)
