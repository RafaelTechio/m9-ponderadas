
# Ponderada Semana 9
Rafael Techio

## Passos

Criar main.go:

![image](https://github.com/user-attachments/assets/0ec958e8-068f-49bb-bb39-e99f88781a34)

Criar arquivo de teste:

![image](https://github.com/user-attachments/assets/ec9da067-6fdd-48a1-9a64-f4b0c08919fd)

Alterar seu código para:

```js
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 1,
  duration: '5s',
}

export default function () {
  http.get('http://192.168.68.108:3000');
  sleep(1);
}
```

Rodar os testes:

![image](https://github.com/user-attachments/assets/e47b9d9b-9b15-4762-b015-ac296ace4b50)

### Load Test

Alterar o código adicionando stages:

```js

import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    stages: [
        {
            duration: '10s',
            target: 100
        },
        {
            duration: '30s',
            target: 100
        },
        {
            duration: '10s',
            target: 0
        }
    ]
}

export default function () {
    http.get('http://192.168.68.108:3000');
    sleep(1);
}
```

Resultado:
![image](https://github.com/user-attachments/assets/d2d1bc83-67f3-493e-b402-b6cfd2a2493e)

## Stress Test

Alterar o código para:

```js
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    stages: [
        {
            duration: '10s',
            target: 200
        },
        {
            duration: '30s',
            target: 200
        },
        {
            duration: '10s',
            target: 0
        }
    ]
}

export default function () {
    http.get('http://192.168.68.108:3000');
    sleep(1);
}
```

Resultado:
![image](https://github.com/user-attachments/assets/720d4300-cf02-497a-91f6-03449581f212)

## Spike Test

Alterar o código para:

```js
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    stages: [
        {
            duration: '1m',
            target: 10000
        },
        {
            duration: '30s',
            target: 0
        }
    ]
}

export default function () {
    http.get('http://192.168.68.108:3000');
    sleep(1);
}
```

Resultados:

![image](https://github.com/user-attachments/assets/6c3b948f-1d04-4f55-bcbe-ae2f8845773f)
