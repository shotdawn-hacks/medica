import pandas as pd
import json

class DataProcessor:
    def __init__(self, data_loaded, data_standart, data_standart_sppvr):
        self.data_loaded = data_loaded
        self.data_standart = data_standart
        self.data_standart_sppvr = data_standart_sppvr
        self.result = pd.DataFrame(columns=['Код МКБ-10', 'Назначения',
                                            'Верное назначение', 'Возможно верное назначение',
                                            'Неверное назначение', 'Верное, но не учтенное'])
    
    def process_data(self):
        self.data_loaded['Назначения'] = self.data_loaded['Назначения'].apply(lambda x: x.strip().split('\n'))
        
        for index, row in self.data_loaded.iterrows():
            code = row['Код МКБ-10']
            appointment = row['Назначения']

            matching_row = self.data_standart_sppvr[self.data_standart_sppvr['Код болезни'].str.strip() == code.strip()]

            if len(matching_row) == 0:
                matching_row = self.data_standart[self.data_standart['Код болезни'].str.strip() == code.strip()]

            if len(matching_row) == 0:
                continue

            required_appointments = matching_row['Обязательные назначения'].iloc[0]
            optional_appointments = matching_row['Необязательные назначения'].iloc[0]

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
                'Код МКБ-10': code,
                'Назначения': appointment,
                'Верное назначение': correct_appointment,
                'Возможно верное назначение': possible_correct_appointment,
                'Неверное назначение': wrong_appointment,
                'Верное, но не учтенное': missed_appointment
            }
    
    def get_result(self):
        return self.result.to_json(orient='records')