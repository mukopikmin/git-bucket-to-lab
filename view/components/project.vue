<template>
  <div class="card">
    <div class="card-header">Project</div>
    <div class="card-body">
      <div v-if="loading" class="text-center my-2">
        <b-spinner variant="primary"></b-spinner>
      </div>

      <div v-else-if="project">
        <h4 class="card-title">{{ project.path_with_namespace }}</h4>
        <p>{{ project.description }}</p>

        <p v-if="isNoBranches">No branches</p>
        <div v-else>
          <h5>Branches</h5>
          <ul>
            <li v-for="branch in project.branches" :key="branch.commit.sha">
              {{ branch.name }} ({{ branch.commit.id }})
            </li>
          </ul>
        </div>
      </div>

      <div v-else>
        <p class="text-center my-2">No project</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['project', 'loading'],
  computed: {
    isNoBranches() {
      return this.project.branches.length === 0
    }
  }
}
</script>

<style></style>
