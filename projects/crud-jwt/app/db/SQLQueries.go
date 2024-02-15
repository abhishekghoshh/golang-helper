package db

const FIND_ALL_ACCOUNT_DETAILS_QUERY = "select user_id,name, address, phone_no, email,date_of_birth, opening_date,last_active_time, status from accounts"
const FIND_ACCOUNT_BY_ID_QUERY = `select user_id,name, address, phone_no, email,date_of_birth, opening_date,last_active_time, status from accounts where user_id = :user_id`
const FIND_USER_BY_ID_QUERY = `select user_id,user_name, password, user_type from users where user_id = :user_id`
