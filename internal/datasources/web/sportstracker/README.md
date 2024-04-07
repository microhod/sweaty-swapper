# Sports Tracker API

```sh
# workouts
curl -H "Sttauthorization:$TOKEN" "https://api.sports-tracker.com/apiserver/v1/workouts?sortonst=true&limit=10&offset=0"

# workout with images
curl -H "Sttauthorization:$TOKEN" "https://api.sports-tracker.com/apiserver/v1/workouts/<workout_id>/combined"

# export GPX
curl "https://api.sports-tracker.com/apiserver/v1/workout/exportGpx/<workout_id>?token=$TOKEN"

# image
curl "https://api.sports-tracker.com/apiserver/v1/image/<image_id>.jpg"
# scaled image
curl "https://api.sports-tracker.com/apiserver/v1/image/scale/<image_id>.jpg?width=200&height=2000&fit=true"
```
