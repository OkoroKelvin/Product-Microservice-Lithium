# POSTMAN DOCUMENTATION FOR GRPC ENDPOINTS

Create Product
{
    "product": {
        "name": "Professional Photography Course 108",
        "description": "Complete digital photography course with downloadable materials",
        "price": 99.99,
        "type": "DIGITAL",
        "digital_products_meta": {
            "download_link": "https://example.com/downloads/photo-courses22",
            "file_size": 10
        }
    }
}

Get Product By ID
{
    "id": "6a08e36e-29c6-4524-98fc-507f60f08cf3"
}

Update Product
{
    "product": {
        "id": "6a08e36e-29c6-4524-98fc-507f60f08cf3",
        "name": "Professional Photography Course 108 Updated",
        "description": "Complete digital photography course with downloadable materials",
        "price": 99.99,
        "type": "DIGITAL",
        "digital_products_meta": {
            "download_link": "https://example.com/downloads/photo-courses22",
            "file_size": 15
        }
    }
}

Delete Product
{
    "id": "6a08e36e-29c6-4524-98fc-507f60f08cf3"
}

List Products
{
    "pagination_query": {
        "page": "1",
        "page_size": "10"
    }
}

Create Subscription

{
  "subscription": {
    "product_id": "6a08e36e-29c6-4524-98fc-507f60f08cf3",
    "plan_name": "premium",
    "duration": 30,
    "price": 99.99
  }
}

Get Subscription By ID
{
 "id": "cd2d4c61-aef6-4bb2-a646-8c4cb0c70843"
 }


 Delete Subscription
 {
    "id": "cd2d4c61-aef6-4bb2-a646-8c4cb0c70843"
 }

 List Subscriptions
 {
    "pagination_query": {
        "page": "1",
        "page_size": "10"
    }
 }

 Update Subscription
{
  "subscription": {
     "id": "cd2d4c61-aef6-4bb2-a646-8c4cb0c70843",
    "product_id": "6a08e36e-29c6-4524-98fc-507f60f08cf3",
    "plan_name": "vip",
    "duration": 30,
    "price": 99.99
  }
}
