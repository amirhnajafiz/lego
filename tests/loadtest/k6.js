import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
    iterations: 1,
    // stages: [
    //     { target: 20, duration: '1m' },
    //     { target: 15, duration: '1m' },
    //     { target: 0, duration: '1m' },
    // ],
    thresholds: {
        requests: ['count < 10'],
    },
};

export default function () {
    const url = "http://localhost:8080";

    // post stage
    const presponse = http.post(url+"/v1/new?key=ktest&value=vtest");
    let pcheckRes = check(presponse, {
        'status is 200': (r) => r.status === 200,
    });
    console.log(pcheckRes)

    sleep(1);

    // get stage
    const gresponse = http.post(url+"/v1/get?key=ktest");
    let gcheckRes = check(gresponse, {
        'status is 200': (r) => r.status === 200,
    });
    console.log(gcheckRes)

    sleep(1);

    // delete stage
    const dresponse = http.post(url+"/v1/del?key=ktest");
    let dcheckRes = check(dresponse, {
        'status is 200': (r) => r.status === 200,
    });
    console.log(dcheckRes)

    sleep(1);
}
