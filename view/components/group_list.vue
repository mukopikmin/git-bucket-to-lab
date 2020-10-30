<template>
  <div>
    <div class="card">
      <table class="table card-table table-hover">
        <thead class="card-header">
          <tr>
            <td>Authorized group</td>
            <td>GitBucket</td>
            <td>GitLab</td>
          </tr>
        </thead>
        <tbody>
          <tr v-if="loading" class="text-center my-5">
            <td colspan="3">
              <b-spinner class="my-3" variant="primary"></b-spinner>
            </td>
          </tr>
          <tr v-for="group in gitbucketGroups" v-else :key="group.login">
            <td>{{ group.login }}</td>
            <td>
              <b-icon-check-circle-fill variant="success" class="ml-1" />
            </td>
            <td>
              <b-icon-check-circle-fill
                v-if="gitlabGroups.map((g) => g.path).includes(group.login)"
                variant="success"
                class="ml-1"
              />
              <b-icon-x v-else variant="danger"></b-icon-x>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  props: {
    loading: Boolean
  },
  computed: {
    ...mapState(['gitbucketGroups', 'gitlabGroups', 'gitlabUrl'])
  }
}
</script>

<style scoped>
.table {
  margin: 0;
}
thead td {
  border-top: none;
}
.icon {
  text-align: center;
}
</style>
