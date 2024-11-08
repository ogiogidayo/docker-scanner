FROM public.ecr.aws/lambda/python:3.8

COPY app.py ${LAMBDA_TASK_ROOT}

RUN npm install -s sample1 sample2

CMD [ "app.handler" ]