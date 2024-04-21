from ultralytics import YOLO

model = YOLO("models-nn/video-model.pt")

def predict_video(source, save, show, project):
    results = model.track(source=source, save=save, show=show, project=project)
    print(results)

def main():
    print("hi")

if __name__ == "__main__":
    main()

predict_video('/home/anna/Work/hackaton/backend/records/test.mp4', True, False, './cache')