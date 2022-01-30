## Currency Converter
###Problem
An API that provides a small but well-thought service converting Euro to US Dollar and vice-versa. That API should only be accessible if you are in the possession of an API-KEY.

###dependent service
To get rates, you can use the free plan of this API: https://fixer.io

###end points
1. localhost:8000/v1/currency_converter/api_key
2. localhost:8000/v1/currency_converter/currency/convert?base=EUR&amount=100
3. localhost:8000/v1/currency_converter/currency/convert?base=USD&amount=100

###sample curls
1. Generate API key
``curl localhost:8000/v1/currency_converter/api_key``
   1. response:
      1. `{"success":true,"api_key":"XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa"}`
2. Convert Euro to Dollar
   1. `localhost:8000/v1/currency_converter/currency/convert?base=EUR&amount=100`
      1. ``{
         "success": true,
         "converted_amount": 111.47890000000001
         }``
3. Convert Dollar to Euro
   1. `localhost:8000/v1/currency_converter/currency/convert?base=USD&amount=100`
      1. `{
         "success": true,
         "converted_amount": 89.7029933888894
         }`
## breaks taken
1. 2022-01-30 19:36 +0530 88ef33359b890829cbdd21f98fea6e38cdb3432c 30 min break
2. 2022-01-30 21:21 +0530 3b8fe805d6edea11536363802d12188e64b7edcc 30 min break 
3. 2022-01-30 22:36 +0530 9a3ed31b60a75d2af72c9bdbf183b799daac5f5d 15 min break
 