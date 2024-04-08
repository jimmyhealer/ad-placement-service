import http from 'k6/http'
import { check, sleep } from 'k6'

export let options = {
    stages: [
        { duration: '5s', target: 10000 },
        { duration: '10s', target: 10000 },
        { duration: '5s', target: 0 },
    ],
    thresholds: {
        http_req_duration: [
            'p(95)<500',
            'p(99)<1500'
        ],
    },
}

export default function () {
    const params = {
        age: 24,
        gender: "F",
        country: "TW",
        platform: "ios",
    }

    const paramsStr = Object.keys(params).map(key => key + '=' + params[key]).join('&')    

    // const response = http.get(`http://localhost:8080/api/v1/ads?${paramsStr}`)
    const response = http.get(`http://localhost:8080/ping`)

    check(response, {
        'is status 200': (r) => r.status === 200,
    })

    sleep(1)
}
