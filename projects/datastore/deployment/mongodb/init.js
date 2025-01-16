db = db.getSiblingDB('mongodb'); // Replace with your database name

// Create unique indexes
db.users.createIndex({ userName: 1 }, { unique: true });
db.users.createIndex({ email: 1 }, { unique: true });