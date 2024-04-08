import http from 'k6/http'
import { check, sleep } from 'k6'

export default function () {
  const params = {
    age: 24,
    gender: "F",
    country: "TW",
    platform: "ios",
  }

  const response = http.get(`http://localhost:8080/api/v1/ads`)

  check(response, {
    'is status 200': (r) => r.status === 200,
  })
}