
<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-white">
    <!-- Header -->
    <div class="relative overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-r from-blue-500/10 via-purple-500/10 to-pink-500/10"></div>
      <div class="relative max-w-7xl mx-auto px-6 py-20">
        <div class="text-center space-y-6">
 <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-bold bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">
            Mes Projets
          </h1>
          <p class="text-xl text-slate-400 max-w-2xl mx-auto">
            Découvrez une sélection de mes réalisations, du frontend au backend, en passant par le mobile
          </p>
        </div>

        <!-- Filtres -->
        <div class="flex flex-wrap justify-center gap-3 mt-12">
          <button
            v-for="cat in categories"
            :key="cat.id"
            @click="filter = cat.id"
            :class="[
              'flex items-center gap-2 px-5 py-2.5 rounded-full font-medium transition-all duration-300',
              filter === cat.id
                ? 'bg-gradient-to-r from-blue-500 to-purple-500 text-white shadow-lg shadow-blue-500/50'
                : 'bg-slate-800/50 text-slate-400 hover:bg-slate-800 hover:text-white border border-slate-700'
            ]"
          >
            <component :is="cat.icon" class="w-4 h-4" />
            {{ cat.label }}
          </button>
        </div>
      </div>
    </div>

    <!-- Grille de projets -->
    <div class="max-w-7xl mx-auto px-6 pb-20">
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="project in filteredProjects"
          :key="project.id"
          @mouseenter="hoveredId = project.id"
          @mouseleave="hoveredId = null"
          @click="selectedProject = project"
          class="group relative bg-slate-900/50 rounded-2xl overflow-hidden border border-slate-800 hover:border-slate-700 transition-all duration-300 cursor-pointer hover:scale-[1.02] hover:shadow-2xl hover:shadow-blue-500/10"
        >
          <!-- Image avec overlay -->
          <div class="relative h-48 overflow-hidden">
            <img
              :src="project.image"
              :alt="project.title"
              class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
            />
            <div :class="['absolute inset-0 bg-gradient-to-t', project.color, 'opacity-0 group-hover:opacity-20 transition-opacity duration-300']"></div>
            <div class="absolute inset-0 bg-gradient-to-t from-slate-900 via-slate-900/50 to-transparent"></div>
            
            <!-- Badge catégorie -->
            <div class="absolute top-3 right-3">
              <span class="px-3 py-1 text-xs font-mono bg-slate-950/80 backdrop-blur-sm border border-slate-700 rounded-full text-slate-300">
                {{ project.category }}
              </span>
            </div>
          </div>

          <!-- Contenu -->
          <div class="p-6 space-y-4">
            <h3 class="text-xl font-bold group-hover:text-transparent group-hover:bg-gradient-to-r group-hover:from-blue-400 group-hover:to-purple-400 group-hover:bg-clip-text transition-all duration-300">
              {{ project.title }}
            </h3>
            
            <p class="text-sm text-slate-400 line-clamp-2">
              {{ project.description }}
            </p>

            <!-- Tags -->
            <div class="flex flex-wrap gap-2">
              <span
                v-for="(tag, i) in project.tags.slice(0, 3)"
                :key="i"
                class="px-2 py-1 text-xs font-mono bg-slate-800/50 border border-slate-700 rounded text-slate-400"
              >
                {{ tag }}
              </span>
              <span v-if="project.tags.length > 3" class="px-2 py-1 text-xs font-mono text-slate-500">
                +{{ project.tags.length - 3 }}
              </span>
            </div>

            <!-- Actions rapides -->
            <div class="flex gap-2 pt-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
              <a
                :href="project.link"
                @click.stop
                target="_blank"
                rel="noopener noreferrer"
                class="flex-1 flex items-center justify-center gap-2 px-3 py-2 bg-blue-500/10 hover:bg-blue-500/20 border border-blue-500/30 rounded-lg text-sm text-blue-400 transition-colors"
              >
                <ExternalLinkIcon class="w-4 h-4" />
                Voir
              </a>
            </div>
          </div>
        </div>
      </div>

      <!-- Message si aucun projet -->
      <div v-if="filteredProjects.length === 0" class="text-center py-20">
        <p class="text-slate-500 text-lg">Aucun projet dans cette catégorie</p>
      </div>
    </div>

    <!-- Modal détails -->
    <div
      v-if="selectedProject"
      class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-6"
      @click="selectedProject = null"
    >
      <div
        class="bg-slate-900 rounded-2xl max-w-4xl w-full max-h-[90vh] overflow-y-auto border border-slate-800"
        @click.stop
      >
        <!-- Image header -->
        <div class="relative h-64 overflow-hidden">
          <img
            :src="selectedProject.image"
            :alt="selectedProject.title"
            class="w-full h-full object-cover"
          />
          <div :class="['absolute inset-0 bg-gradient-to-t', selectedProject.color, 'opacity-30']"></div>
          <div class="absolute inset-0 bg-gradient-to-t from-slate-900 via-slate-900/50 to-transparent"></div>
          
          <button
            @click="selectedProject = null"
            class="absolute top-4 right-4 w-10 h-10 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm hover:bg-slate-950 border border-slate-700 rounded-full text-white transition-colors"
          >
            ✕
          </button>
        </div>

        <!-- Contenu détaillé -->
        <div class="p-8 space-y-6">
          <div>
            <h2 class="text-3xl font-bold mb-2">{{ selectedProject.title }}</h2>
            <p class="text-slate-400">{{ selectedProject.longDescription }}</p>
          </div>

          <!-- Tous les tags -->
          <div>
            <h3 class="text-sm font-semibold text-slate-500 uppercase tracking-wider mb-3">Technologies</h3>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="(tag, i) in selectedProject.tags"
                :key="i"
                class="px-3 py-1.5 text-sm font-mono bg-slate-800/50 border border-slate-700 rounded-lg text-slate-300"
              >
                {{ tag }}
              </span>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex gap-3 pt-4">
            <a
              :href="selectedProject.link"
              target="_blank"
              rel="noopener noreferrer"
              class="flex-1 flex items-center justify-center gap-2 px-6 py-3 bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 rounded-lg font-medium transition-all duration-300 shadow-lg shadow-blue-500/30"
            >
              <ExternalLinkIcon class="w-5 h-5" />
              Voir le projet
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';

// Icônes SVG
const ExternalLinkIcon = {
  template: '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" /></svg>'
};

const GithubIcon = {
  template: '<svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>'
};

const LayersIcon = {
  template: '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7l8-4 8 4-8 4-8-4zm0 5l8 4 8-4m-8 4l-8 4 8 4 8-4-8-4z"/></svg>'
};

const CodeIcon = {
  template: '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>'
};

const SparklesIcon = {
  template: '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"/></svg>'
};

// State
const hoveredId = ref(null);
const selectedProject = ref(null);
const filter = ref('all');

// Données
const projects = ref([

  {
    id: 1,
    title: "ösleya  b.",
    description: "Création de vêtements personnalisés",
    longDescription: "Site vitrine pour la société ösleya b. spécialisée dans la conception, la personnalisation et la fabrication de vestes sur-mesures.",
    image: "https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=800&h=600&fit=crop",
    tags: ["Vue.js", "Email JS", "Tailwind"],
    category: "frontend",
    color: "from-purple-500 to-pink-500",
    link: "https://osleyab.fr",
  },
 // Copié collé d'autres projets ici
]);

const categories = ref([
  { id: 'all', label: 'Tous', icon: LayersIcon },
  { id: 'frontend', label: 'Frontend', icon: CodeIcon },
  { id: 'backend', label: 'Backend', icon: SparklesIcon },
  { id: 'fullstack', label: 'Fullstack', icon: LayersIcon },
  { id: 'mobile', label: 'Mobile', icon: CodeIcon }
]);

// Computed
const filteredProjects = computed(() => {
  return filter.value === 'all' 
    ? projects.value 
    : projects.value.filter(p => p.category === filter.value);
});
</script>

<style scoped>
/* Styles personnalisés si nécessaire */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>