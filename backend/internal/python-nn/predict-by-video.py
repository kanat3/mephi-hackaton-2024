from ultralytics import YOLO
import sys

model = YOLO("./internal/python-nn/models-nn/video-model.pt")

def predict_video(source, save, show, project):
    results = model.track(source=source, save=save, show=show, project=project)
    print(results)

def main():
    print("Run predict by video")
    predict_video(sys.argv[1], bool(sys.argv[2]), False, sys.argv[3])

if __name__ == "__main__":
    main()
