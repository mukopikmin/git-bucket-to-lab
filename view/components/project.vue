<template>
  <b-card no-body>
    <b-card-header>Project</b-card-header>
    <b-card-body>
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="project">
        <b-card-title>
          <b-icon-lock v-if="isPrivate" class="mr-1" />
          <b-icon-bookmarks v-else class="mr-1" />
          {{ project.namespace.path }} / {{ project.path }}
        </b-card-title>
        <b-card-text>{{ project.description }}</b-card-text>
      </div>

      <div v-else>
        <p class="text-center my-2">No project</p>
      </div>
    </b-card-body>

    <div v-if="project">
      <h5 class="list-title">Branches</h5>
      <b-list-group flush class="border-top border-bottom">
        <b-list-group-item v-if="isNoBranches">No branches</b-list-group-item>
        <b-list-group-item
          v-for="branch in project.branches"
          :key="branch.commit.sha"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <Branch class="branch" :name="branch.name" :sha="branch.commit.id" />
        </b-list-group-item>
      </b-list-group>

      <h5 class="list-title mt-4">Tags</h5>
      <b-list-group flush class="border-top">
        <b-list-group-item v-if="isNoTags">No tags</b-list-group-item>
        <b-list-group-item
          v-for="tag in project.tags"
          :key="tag.commit.id"
          class="d-flex justify-content-between align-items-center text-align-left"
        >
          <Branch class="branch" :name="tag.name" :sha="tag.commit.id" />
        </b-list-group-item>
      </b-list-group>
    </div>

    <template v-if="project" v-slot:footer>
      <small class="text-muted">
        <a :href="project.web_url" target="_blank">
          <b-icon-box-arrow-up-right class="mr-1" />
          <span>Open project</span>
        </a>
      </small>
    </template>
  </b-card>
</template>

<script>
import Branch from '@/components/branch'

export default {
  components: {
    Branch
  },
  props: {
    project: {
      type: Object,
      required: false,
      default: null
    },
    loading: Boolean
  },
  computed: {
    isNoBranches() {
      return this.project.branches.length === 0
    },
    isNoTags() {
      return this.project.tags.length === 0
    },
    isPrivate() {
      return this.project.visibility === 'private'
    }
  }
}
</script>

<style scoped>
.branch {
  width: 100%;
}
.list-title {
  padding-left: 20px;
}
.border-top {
  border-top: 1px solid rgba(0, 0, 0, 0.125);
}
.border-bottom {
  border-bottom: 1px solid rgba(0, 0, 0, 0.125);
}
</style>
