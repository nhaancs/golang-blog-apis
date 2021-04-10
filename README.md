# ecommerce

## Questions
- A Việt ơi ví dụ trong tầng biz e validate thấy email invalid và muốn tầng transport trả về status 400 cho user, còn khi insert xuống db lỗi e muốn trả về 500 thì làm như thế nào ak a

- 2 component cùng layer co the goi lan nhau, layer o tren co the gọi xuống bên dưới, Có khi nào 2 biz gọi nhau?

GIN_MODE=debug DB_CONN_STR="food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" go run .

https://jamstackvietnam.com/bai-viet/72-yeu-to-khong-the-thieu-khi-thiet-ke-website-thuong-mai-dien-tu 

## Entities
- Product (featured, recommanded for user)
- User
- Setting (Logo, SEO, lat, long)
- Slide (Unique value proposition)
- Blog category (name desc)
- Blog post
- Contact
- Payment
- Order
- Policy

- Favorite
- Product category
- File
- Collection (product collections, CTA) 
- Banner (text, link to a post, collection, product)