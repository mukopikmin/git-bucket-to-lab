<template>
  <b-card no-body header="Issues">
    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item
          v-for="issue in issues"
          :key="`gitlab-issue-${issue.id}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="issue.state == 'opened'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ issue.iid }} {{ issue.title }}
          </span>
          <b-badge
            v-if="issue.comments && issue.comments.length > 0"
            variant="primary"
            pill
            >{{ issue.comments.length }}</b-badge
          >
        </b-list-group-item>
      </b-list-group>
    </div>
  </b-card>
</template>

<script>
export default {
  props: ['issues', 'loading']
}
</script>
