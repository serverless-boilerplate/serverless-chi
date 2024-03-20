FROM public.ecr.aws/lambda/go:1.2024.03.04.11

COPY src go.mod go.sum ${LAMBDA_TASK_ROOT}

RUN go build
