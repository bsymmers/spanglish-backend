import re
import pandas as pd
import numpy as np
from sklearn.preprocessing import LabelEncoder
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.model_selection import train_test_split
from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import accuracy_score, confusion_matrix, classification_report

class DetectLanguageModel:
    def __init__(self):
        self.label_encoder = None
        self.model = None
        self.cv = None
    def build_model(self) -> None:
        """
        Based on https://heartbeat.comet.ml/using-machine-learning-for-language-detection-517fa6e68f22
        """
        df = pd.read_csv("ml/Language Detection.csv")
        X = df["Text"]
        y = df["Language"]

        text_list = []
        for text in X:
            text = re.sub(r'[!@#$(),n"%^*?:;~`0-9]', ' ', text)
            text = re.sub(r'[[]]', ' ', text)
            text = text.lower()
            text_list.append(text)

        self.label_encoder = LabelEncoder()

        y = self.label_encoder.fit_transform(y)
        self.cv = CountVectorizer()
        X = self.cv.fit_transform(text_list).toarray()
        x_train, x_test, y_train, y_test = train_test_split(X, y, test_size = 0.20)
        
        self.model = MultinomialNB()
        self.model.fit(x_train, y_train)

        # y_prediction = model.predict(x_test)
        # accuracy = accuracy_score(y_test, y_prediction)
        # confusion_m = confusion_matrix(y_test, y_prediction)
        # print("The accuracy is :",accuracy)
    def detect_language(self, text : str) -> str:
        x = self.cv.transform([text]).toarray()
        lang = self.model.predict(x)
        lang = self.label_encoder.inverse_transform(lang)
        print("The language is in", lang[0])

# process_language()
# instance = DetectLanguageModel()
# instance.build_model()
# instance.detect_language("Hola. Soy brandon.")
