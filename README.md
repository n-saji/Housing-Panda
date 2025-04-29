# HousingPanda Backend – House Rental Listing Service

A simple backend service where users can add their house listings for rent and fetch all available houses for sublease.

---

## Tech Stack
- **Golang**
- **PostgreSQL**

---

##  API Usage

### Add a New Listing (POST `/listings`)

#### Sample Request Body:
```json
{
  "title": "Fully Furnished Beautiful House",
  "description": "A well maintained house",
  "address": "155 Merrimac Avn, Buffalo, New York - 14236",
  "rent": 1999.99,
  "number_of_bedrooms": 2,
  "number_of_bathrooms": 1,
  "name": "Jhon Mathew",
  "phone_number": "4162393699"
}
```

## Available Endpoints

| Method | Endpoint                    | Description                            |
|--------|-----------------------------|----------------------------------------|
| POST   | `/listings`                 | Create a new house listing             |
| GET    | `/listings`                 | Get all available house listings       |
| GET    | `/listings/:listing_id`     | Get details of a specific listing      |
| GET    | `/listings/user/:user_id`   | Get all listings from a specific user  |
| DELETE | `/listings/:listing_id`     | Delete a listing                       |
| GET    | `/users/:user_id`           | Get user details                       |
| GET    | `/users`                    | Get all users                          |

