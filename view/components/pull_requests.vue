<template>
  <b-card no-body>
    <b-card-header header-tag="nav">
      <span class="title">Pull Requests</span>
      <MigrateButton :action="migratePulls" />
    </b-card-header>

    <div v-if="loading" class="text-center my-2">
      <b-spinner variant="primary"></b-spinner>
    </div>

    <div v-else>
      <b-list-group flush>
        <b-list-group-item
          v-for="pull in pulls"
          :key="`pull-${pull.number}`"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <span>
            <b-icon-check-circle
              v-if="pull.state == 'open'"
              variant="success"
              class="mr-1"
            ></b-icon-check-circle>
            <b-icon-slash-circle
              v-else
              variant="danger"
              class="mr-1"
            ></b-icon-slash-circle>
            #{{ pull.number }} {{ pull.title }}
          </span>
          <b-badge v-if="pull.comments.length > 0" variant="primary" pill>{{
            pull.comments.length
          }}</b-badge>
        </b-list-group-item>
      </b-list-group>
    </div>
  </b-card>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import MigrateButton from '@/components/migrate_button'

export default {
  components: {
    MigrateButton
  },
  props: ['repo', 'pulls', 'loading'],
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  methods: {
    ...mapActions(['setRepo', 'setProject']),
    async migratePulls() {
      const res = await this.$axios.$post(
        `/${this.repo.owner.login}/${this.repo.name}/pulls`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.setRepo(res.repo)
      this.setProject(res.project)
    }
  }
}
</script>

<style scoped>
.card-header {
  padding-top: 8.5px;
  padding-bottom: 8.5px;
}
.title {
  height: 100%;
  vertical-align: middle;
}
.migrate-button {
  float: right;
}
</style>
