<template>
  <b-card no-body header="Merge Requests">
    <b-card-body v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </b-card-body>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item v-if="noMerges">No merge requests</b-list-group-item>
        <b-list-group-item
          v-for="merge in pagedMerges"
          :key="`merge-${merge.iid}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="merge.state == 'opened'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ merge.iid }} {{ merge.title }}
          </span>
          <b-badge
            v-if="merge.comments && merge.comments.length > 0"
            variant="primary"
            pill
            >{{ merge.comments.length }}</b-badge
          >
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
      <small class="text-muted"
        ><b-icon-info-circle class="mr-1" />{{ openCount }} Open
        <b-icon-check2 class="ml-2 mr-1" />{{ closedCount }} Closed</small
      >
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
    merges: {
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
    pagedMerges() {
      return this.merges.slice(
        this.perPage * (this.page - 1),
        this.perPage * this.page
      )
    },
    pageSize() {
      return Math.ceil(this.merges.length / this.perPage)
    },
    paginationEnabled() {
      return this.pageSize > 1
    },
    noMerges() {
      return this.merges.length === 0
    },
    openCount() {
      return this.merges.filter((m) => m.state === 'opened').length
    },
    closedCount() {
      return this.merges.filter((m) => m.state === 'closed').length
    }
  },
  methods: {
    onPageChange(e) {
      this.page = e
    }
  }
}
</script>

<style></style>
