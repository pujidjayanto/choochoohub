# Overview
Choochoohub is a backend system for train booking application. The reason i created choochoohub is to learn about building microservices. And i was thinking what backend application i can copy?. Because i usually book train using KAI (Kereta Api Indonesia) app, so i decied to build a backend system that copy the functionality of KAI application.

# Business flow
To know the business flow, i reflect on my activities on the app when using the app.
1. Sign up using email and password
2. You got the OTP sent to your email
3. Put and verify OTP
4. If verified, you will go to the main menu, where there are several buttons there.
5. There is a profile menu.
6. When you update your profile, and fill necessary informations, you will be upgraded to premium profile.
7. You also can add, remove, update passanger data for your profile.
8. Then there is booking menu.
9. First you need to search schedule. You put search entry to search schedules.
10. You will get list of trains of your selected date. You can pick one.
11. Then you can choose the passengers.
12. Then you can choose seats if available.
13. After all settled, you create a booking.
14. You can pick payment methods. I usually do QRIS.
15. Then you pay. If you already paid, go back to home menu.
16. Then you can see your booked schedules.
17. And also there is My tickets menu.
18. It show all your booked schedules, from past, present, and future.


# Service plan
1. Api Gateway -> as main point. call other services. also generate token.
2. User API -> handle users data.
3. Inventory API -> handle train, and schedules data.
4. Notification client -> handle notification.
5. Booking API -> handle booking.
6. Payment API -> mimic payment service.

# Apis plan
## Api Gateway
1. POST /v1/signup
2. POST /v1/signup/verify-otp
3. POST /v1/signin
4. GET v1/schedules/search
5. GET v1/schedules/:id/detail
6. POST v1/bookings
7. POST v1/payment
8. GET v1/schedules/
9. GET v1/users/:id
10. PUT v1/users/:id

## User Api
1. POST /v1/signup
2. POST /v1/signup/verify-otp
3. POST /v1/signin
4. GET v1/users/:id
5. PUT v1/users/:id

# Techstack
1. Database will be postgresql
2. Language mostly Golang. But i want to try haskell later
3. HTTP and GRPC
4. Kafka
5. others...

# Current
1. As for right now, i only planning for User Api.
2. Since i want to look for new job, i need to do preparation.
3. Therefore, the development will be suspended or i will work on it slowly.


# Todo
1. user-api: create job to expire otp
