from ultralytics import YOLO

model = YOLO("models-nn/video-model.pt")

def predict_video(source, show, save):
    print("hi")
    results = model.track(source, show, save)