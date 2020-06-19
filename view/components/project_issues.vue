<template>
  <b-card no-body header="Issues">
    <b-card-body v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </b-card-body>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item v-if="noIssues">No issuess</b-list-group-item>
        <b-list-group-item
          v-for="issue in pagedIssues"
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
          >
            {{ issue.comments.length }}
          </b-badge>
        </b-list-group-item>
      </b-list-group>

      <Pagination
        v-if="paginationEnabled"
        :page="page"
        :page-size="pageSize"
        @change="onPageChange"
      />
    </div>

    <template v-slot:footer>
      <small class="text-muted">
        <b-icon-info-circle class="mr-1" />
        {{ openCount }} Open
        <b-icon-check2 class="ml-2 mr-1" />
        {{ closedCount }} Closed
      </small>
    </template>
  </b-card>
</template>

<script>
import Pagination from '@/components/pagination'

export default {
  components: {
    Pagination
  },
  props: {
    issues: {
      type: Array,
      required: true
    },
    loading: Boolean
  },
  data() {
    return {
      page: 1,
      perPage: process.env.pageSize
    }
  },
  computed: {
    pagedIssues() {
      return this.issues.slice(
        this.perPage * (this.page - 1),
        this.perPage * this.page
      )
    },
    pageSize() {
      return Math.ceil(this.issues.length / this.perPage)
    },
    paginationEnabled() {
      return this.pageSize > 1
    },
    noIssues() {
      return this.issues.length === 0
    },
    openCount() {
      return this.issues.filter((i) => i.state === 'opened').length
    },
    closedCount() {
      return this.issues.filter((i) => i.state === 'closed').length
    }
  },
  methods: {
    onPageChange(e) {
      this.page = e
    }
  }
}
</script>
