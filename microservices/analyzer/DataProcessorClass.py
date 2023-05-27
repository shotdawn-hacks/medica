import pandas as pd
import json

class DataProcessor:
    def __init__(self, data_loaded, data_standart):
        self.data_loaded = data_loaded
        self.data_standart = data_standart
        self.result = []
    
    def process_data(self):
        for index, row in self.data_loaded.iterrows():
            code = row['Код МКБ-10']
            appointment = row['Назначения']
            
            standart_row = self.data_standart[self.data_standart['Код болезни'].str.strip() == code.strip()]
            
            if standart_row.empty:
                encoded_appointment = [0]
            else:
                required_appointments = standart_row['Обязательные назначения'].iloc[0]
                optional_appointments = standart_row['Необязательные назначения'].iloc[0]
            
                encoded_appointment = []
            
                for app in appointment:
                    if app in required_appointments:
                        encoded_appointment.append(1)
                    elif app in optional_appointments:
                        encoded_appointment.append(2)
                    else:
                        encoded_appointment.append(3)
            
            result_item = {
                'Код МКБ-10': code,
                'Назначения': appointment,
                'Закодированные назначения': encoded_appointment
            }
            
            self.result.append(result_item)
    
    def get_result(self):
        return json.dumps(self.result)
