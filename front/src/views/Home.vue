<template lang="pug">
  .home
    span(v-if="!loading && todo")
      | {{ todo.getTimestamp().toDate() }}
      | {{ todo.getText() }}
    span(v-else)
      | loading...

    .error(v-if="error")
      | {{ error }}
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { grpc } from "grpc-web-client";
import { GetLatestRequest, GetLatestResponse, Todo } from "@/pb/todo_pb";
import { TodoService } from "@/pb/todo_pb_service";

@Component({})
export default class Home extends Vue {
  error: string | null;
  loading: boolean;
  todo: Todo | null;

  constructor() {
    super();

    this.error = null;
    this.loading = true;
    this.todo = null;
  }

  mounted() {
    const req = new GetLatestRequest();

    grpc.unary(TodoService.GetLatest, {
      request: req,
      host: "http://127.0.0.1:7777",
      onEnd: res => {
        const { status, statusMessage, headers, message, trailers } = res;
        if (status === grpc.Code.OK && message) {
          this.error = null;

          const resObj: GetLatestResponse = message as GetLatestResponse;
          const t = resObj.getTodo();
          if (t !== undefined) {
            this.todo = t;
          }

          this.loading = false;
        } else {
          this.loading = false;
          this.error = res.statusMessage;
        }
      },
    });
  }
}
</script>

<style lang="less">
.error {
  color: #f00;
}
</style>
