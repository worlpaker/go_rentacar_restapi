db = db.getSiblingDB("GODB");

db.locations.insertMany([
    { id: 1, name: "Istanbul", active: true },
    { id: 2, name: "New York City", active: false },
    { id: 3, name: "Paris", active: true },
    { id: 4, name: "London", active: true },
]);

db.offices.insertMany([
    { id: 1, location_id: 1, opening_hour: "08.00", closing_hour: "18.00", working_days: [1, 2, 3, 4, 5] },
    { id: 2, location_id: 3, opening_hour: "12.00", closing_hour: "15.00", working_days: [6, 7] },
    { id: 3, location_id: 2, opening_hour: "00.00", closing_hour: "23.59", working_days: [1, 3, 5] },
    { id: 4, location_id: 4, opening_hour: "06.30", closing_hour: "21.00", working_days: [1, 2, 7] },
]);

db.cars.insertMany([
    { id: 1, fuel: "Petrol", transmission: "Manual", vendor: "Ford", name: "Ford Focus", office_id: 1, reserved: false, reserved_by: {} },
    { id: 2, fuel: "Petrol", transmission: "Automatic", vendor: "Mercedes", name: "Mercedes E200", office_id: 4, reserved: false, reserved_by: {} },
    { id: 3, fuel: "Diesel", transmission: "Automatic", vendor: "Bmw", name: "BMW 5.20", office_id: 2, reserved: false, reserved_by: {} },
    { id: 4, fuel: "Electric", transmission: "Hybrid", vendor: "Volkswagen", name: "Volkswagen Golf", office_id: 3, reserved: false, reserved_by: {} },
]);