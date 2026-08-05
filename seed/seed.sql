BEGIN;

-----------------------------------------------------
-- Hospital
-----------------------------------------------------

INSERT INTO hospitals (id, name, api_url)
VALUES
    (1, 'Bangkok Hospital', 'https://his.bangkok.local'),
    (2, 'Chiang Mai Hospital', 'https://his.chiangmai.local');

-----------------------------------------------------
-- Patient
-----------------------------------------------------

INSERT INTO patients (
    id,
    date_of_birth,
    national_id,
    gender
)
VALUES
    (1, '1990-01-15', '1100000000001', 'Male'),
    (2, '1992-04-10', '1100000000002', 'Female'),
    (3, '1988-12-01', '1100000000003', 'Male'),
    (4, '2001-08-20', '1100000000004', 'Female'),
    (5, '1995-06-30', '1100000000005', 'Male'),
    (6, '1985-03-14', NULL, 'Female'),
    (7, '1998-09-09', NULL, 'Male'),
    (8, '2000-11-22', '1100000000008', 'Female');

-----------------------------------------------------
-- Patient Hospital
-----------------------------------------------------

INSERT INTO patient_hospitals (
    patient_id,
    hospital_id,
    first_name_th,
    middle_name_th,
    last_name_th,
    first_name_en,
    middle_name_en,
    last_name_en,
    passport_id,
    phone_number,
    email,
    patient_hn
)
VALUES

-- Bangkok Hospital

(
    1,
    1,
    'สมชาย',
    '',
    'ใจดี',
    'Somchai',
    '',
    'Jaidee',
    NULL,
    '0811111111',
    'somchai@gmail.com',
    'BK000001'
),

(
    2,
    1,
    'สมหญิง',
    '',
    'สุขใจ',
    'Somying',
    '',
    'Sukjai',
    NULL,
    '0822222222',
    'somying@gmail.com',
    'BK000002'
),

(
    3,
    1,
    'วีระ',
    '',
    'แสงทอง',
    'Weera',
    '',
    'Saengthong',
    'AA123456',
    '0833333333',
    'weera@gmail.com',
    'BK000003'
),

(
    4,
    1,
    'อรทัย',
    '',
    'บุญมี',
    'Orathai',
    '',
    'Boonmee',
    NULL,
    '0844444444',
    'orathai@gmail.com',
    'BK000004'
),

(
    5,
    1,
    'ธนา',
    '',
    'มีสุข',
    'Thana',
    '',
    'Meesuk',
    NULL,
    '0855555555',
    'thana@gmail.com',
    'BK000005'
),

-- Chiang Mai Hospital

(
    6,
    2,
    'กัญญา',
    '',
    'คำดี',
    'Kanya',
    '',
    'Khamdee',
    NULL,
    '0866666666',
    'kanya@gmail.com',
    'CM000001'
),

(
    7,
    2,
    'ปกรณ์',
    '',
    'วงศ์ดี',
    'Pakorn',
    '',
    'Wongdee',
    'BB987654',
    '0877777777',
    'pakorn@gmail.com',
    'CM000002'
),

(
    8,
    2,
    'สุชาดา',
    '',
    'บุญช่วย',
    'Suchada',
    '',
    'Boonchuay',
    NULL,
    '0888888888',
    'suchada@gmail.com',
    'CM000003'
);

COMMIT;

SELECT setval(
               pg_get_serial_sequence('hospitals', 'id'),
               (SELECT MAX(id) FROM hospitals)
       );

SELECT setval(
               pg_get_serial_sequence('patients', 'id'),
               (SELECT MAX(id) FROM patients)
       );

SELECT setval(
               pg_get_serial_sequence('staffs', 'id'),
               COALESCE((SELECT MAX(id) FROM staffs), 1)
       );