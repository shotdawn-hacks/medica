import pandas as pd
import json
import ast

class DataProcessor:
    def __init__(self, data_standard, data_standard_sppvr):
        self.data_standard = pd.read_csv(data_standard)
        self.data_standard_sppvr = pd.read_csv(data_standard_sppvr)
        self.result = pd.DataFrame(columns=['_id','icd', 'prescription',
                                            'correct_prescription', 'possible_correct_prescription',
                                            'wrong_prescription', 'correct_but_not_accounted_prescription'])
        
        self.data_standard_sppvr['Required_Appointments'] = self.data_standard_sppvr['Required_Appointments'].apply(ast.literal_eval)
        self.data_standard_sppvr['Optional_Appointments'] = self.data_standard_sppvr['Optional_Appointments'].apply(ast.literal_eval)
        self.data_standard['Required_Appointments'] = self.data_standard['Required_Appointments'].apply(ast.literal_eval)
        self.data_standard['Optional_Appointments'] = self.data_standard['Optional_Appointments'].apply(ast.literal_eval)
    
    def load_data(self, json_file):
        data = json_file
        self.data_loaded = pd.DataFrame(data)
        self.data_loaded.columns = ['_id', 'icd', 'prescription']
    
    def process_data(self):
        self.data_loaded['prescription'] = self.data_loaded['prescription'].apply(lambda x: x.strip().split('\n'))
        
        for index, row in self.data_loaded.iterrows():
            _id = row["_id"]
            code = row['icd']
            appointment = row['prescription']

            matching_row = self.data_standard_sppvr[self.data_standard_sppvr['ICD-10_Code'].str.strip() == code.strip()]

            if len(matching_row) == 0:
                matching_row = self.data_standard[self.data_standard['ICD-10_Code'].str.strip() == code.strip()]

            if len(matching_row) == 0:
                continue

            required_appointments = matching_row['Required_Appointments'].iloc[0]
            optional_appointments = matching_row['Optional_Appointments'].iloc[0]

            correct_appointment = []
            possible_correct_appointment = []
            wrong_appointment = []
            missed_appointment = []

            for app in appointment:
                if app.strip() in required_appointments:
                    correct_appointment.append(app)
                elif app.strip() in optional_appointments:
                    possible_correct_appointment.append(app)
                else:
                    wrong_appointment.append(app)

            for req_app in required_appointments:
                if req_app.strip() not in appointment:
                    missed_appointment.append(req_app)

            self.result.loc[index] = {
                "_id": _id,
                'icd': code,
                'prescription': appointment,
                'correct_prescription': correct_appointment,
                'possible_correct_prescription': possible_correct_appointment,
                'wrong_prescription': wrong_appointment,
                'correct_but_not_accounted_prescription': missed_appointment
            }
    
    def get_result(self):
        return self.result.to_json(orient='records', force_ascii=False)