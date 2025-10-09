<template>
  <div>
    <!-- Hero Section avec profil -->
    <div class="relative min-h-screen flex flex-col bg-gradient-to-b from-gray-900 to-gray-800 overflow-hidden">
      <!-- Grille animée de fond -->
      <div class="absolute inset-0 opacity-10">
        <div class="absolute inset-0 grid-background"></div>
      </div>

      <!-- Orbes lumineux -->
      <div class="absolute top-20 left-20 w-96 h-96 bg-blue-500 rounded-full filter blur-3xl opacity-20 animate-pulse"></div>
      <div class="absolute bottom-20 right-20 w-96 h-96 bg-blue-400 rounded-full filter blur-3xl opacity-15 animate-pulse-slow"></div>

      <div class="relative z-10 pt-10 flex-1 flex flex-col items-center justify-center text-white p-4 space-y-8">
        <!-- Image avec style minimaliste moderne -->
        <div class="relative group">
          <!-- Ombre douce -->
          <div class="absolute inset-0 bg-blue-500/20 rounded-full blur-3xl opacity-0 group-hover:opacity-100 transition-opacity duration-700"></div>
          
          <div class="relative">
            <!-- Cadre fin animé -->
            <div class="absolute inset-0 rounded-full">
              <svg class="w-full h-full" viewBox="0 0 100 100">
                <circle 
                  cx="50" 
                  cy="50" 
                  r="48" 
                  fill="none" 
                  stroke="url(#gradient)" 
                  stroke-width="0.3"
                  stroke-dasharray="2 4"
                  class="circle-rotate"
                />
                <defs>
                  <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                    <stop offset="0%" style="stop-color:#3b82f6;stop-opacity:0.4" />
                    <stop offset="50%" style="stop-color:#8b5cf6;stop-opacity:0.6" />
                    <stop offset="100%" style="stop-color:#3b82f6;stop-opacity:0.4" />
                  </linearGradient>
                </defs>
              </svg>
            </div>
            
            <!-- Image circulaire -->
            <div class="relative w-80 h-80 rounded-full overflow-hidden border border-gray-700/30 group-hover:border-blue-500/30 transition-all duration-500">
              <img 
                src="../components/Images/myself.jpeg" 
                alt="Profile" 
                class="w-full h-full object-cover filter grayscale-0 group-hover:scale-105 transition-all duration-700" 
              />
              
              <!-- Overlay subtil -->
              <div class="absolute inset-0 bg-gradient-to-br from-blue-500/0 to-blue-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-700"></div>
            </div>
            
            <!-- Petits accents aux 4 points cardinaux -->
            <div class="accent-dot accent-top"></div>
            <div class="accent-dot accent-right"></div>
            <div class="accent-dot accent-bottom"></div>
            <div class="accent-dot accent-left"></div>
          </div>
        </div>

        <div class="card-body max-w-3xl">
          <!-- Badge Industrie 4.0 -->
          <div class="inline-block px-4 py-2 bg-blue-500/20 border border-blue-400/30 rounded-full mb-4 animate-fade-in">
            <span class="text-blue-400 font-mono text-sm tracking-wider">INDUSTRIE 4.0</span>
          </div>

          <h5 class="tracking-wide card-title manrope mb-2 pt-10 text-3xl mx-auto text-center font-bold bg-gradient-to-r from-blue-400 to-blue-200 bg-clip-text text-transparent animate-fade-in-up">
            KABRÉ SYLVAIN
          </h5>
          
          <!-- Ligne décorative -->
          <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-8 animate-fade-in-up"></div>

          <p class="text-xl pt-5 text-gray-200 leading-relaxed animate-fade-in-up-delay">
            Ma passion pour la <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold border-b-2 border-transparent hover:border-blue-400" href="#">technologie</a> va de pair avec ma recherche constante de défis intellectuels.
            Mon parcours <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold border-b-2 border-transparent hover:border-blue-400" href="#">professionnel</a> s'est construit autour de la résolution de problèmes complexes, une démarche qui a façonné mes compétences. 
            Si vous partagez ces aspirations ou souhaitez échanger sur des sujets technologiques, n'hésitez pas à me <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold border-b-2 border-transparent hover:border-blue-400" href="#">contacter.</a>
          </p>

          <!-- Boutons CTA -->
          <div class="flex gap-4 justify-center mt-8 animate-fade-in-up-delay">
            <button class="px-6 py-3 bg-blue-500 hover:bg-blue-600 rounded-lg font-semibold transition-all duration-300 hover:scale-105 hover:shadow-lg hover:shadow-blue-500/50">
              Télécharger mon CV
            </button>
            <button class="px-6 py-3 bg-transparent border-2 border-blue-400 hover:bg-blue-400/10 rounded-lg font-semibold transition-all duration-300 hover:scale-105">
              Me contacter
            </button>
          </div>
        </div>
      </div>

      <!-- Indicateur de scroll -->
      <div class="absolute bottom-2 left-1/2 transform -translate-x-1/2 animate-bounce">
        <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
      </div>
    </div>

    <!-- Statistiques animées -->
    <div class="relative py-20 bg-gray-900">
      <div class="max-w-6xl mx-auto px-4">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedYears }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Années d'expérience</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedProjects }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Projets réalisés</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedTech }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Technologies</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedClients }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Clients satisfaits</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Section Compétences -->
    <div class="relative py-20 bg-gray-800">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white">Expertise Technique</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12"></div>

        <div class="grid md:grid-cols-3 gap-6">
          <!-- Frontend -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">⚡</div>
            <h3 class="text-xl font-semibold text-white mb-4">Frontend</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">V</span>
                </div>
                <span class="text-gray-300">Vue.js</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">R</span>
                </div>
                <span class="text-gray-300">React</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">T</span>
                </div>
                <span class="text-gray-300">Tailwind CSS</span>
              </div>
            </div>
          </div>

          <!-- Backend -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">🔧</div>
            <h3 class="text-xl font-semibold text-white mb-4">Backend</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">P</span>
                </div>
                <span class="text-gray-300">Python</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">N</span>
                </div>
                <span class="text-gray-300">Node.js</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">D</span>
                </div>
                <span class="text-gray-300">Databases</span>
              </div>
            </div>
          </div>

          <!-- DevOps -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">🚀</div>
            <h3 class="text-xl font-semibold text-white mb-4">DevOps & Tools</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">G</span>
                </div>
                <span class="text-gray-300">Git</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">D</span>
                </div>
                <span class="text-gray-300">Docker</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">C</span>
                </div>
                <span class="text-gray-300">CI/CD</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Timeline Parcours -->
    <div class="relative py-20 bg-gray-900">
      <div class="max-w-4xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white">Parcours</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12"></div>

        <div class="relative">
          <!-- Ligne verticale -->
          <div class="absolute left-1/2 transform -translate-x-1/2 h-full w-0.5 bg-blue-500/30"></div>

          <!-- Expérience 1 -->
          <div class="relative mb-12 flex items-center">
            <div class="w-1/2 pr-8 text-right">
              <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300">
                <h3 class="text-xl font-semibold text-white mb-2">Ingénieur Tests</h3>
                <p class="text-blue-400 mb-2">2020 - Présent</p>
                <p class="text-gray-400 text-sm">Développement d'applications de test pour l'industrie automobile</p>
              </div>
            </div>
            <div class="absolute left-1/2 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-900"></div>
            <div class="w-1/2"></div>
          </div>

          <!-- Expérience 2 -->
          <div class="relative mb-12 flex items-center">
            <div class="w-1/2"></div>
            <div class="absolute left-1/2 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-900"></div>
            <div class="w-1/2 pl-8">
              <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300">
                <h3 class="text-xl font-semibold text-white mb-2">Développeur Full-Stack</h3>
                <p class="text-blue-400 mb-2">2018 - 2020</p>
                <p class="text-gray-400 text-sm">Création d'applications web modernes et performantes</p>
              </div>
            </div>
          </div>

          <!-- Formation -->
          <div class="relative mb-12 flex items-center">
            <div class="w-1/2 pr-8 text-right">
              <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300">
                <h3 class="text-xl font-semibold text-white mb-2">Formation Ingénieur</h3>
                <p class="text-blue-400 mb-2">2015 - 2018</p>
                <p class="text-gray-400 text-sm">Spécialisation en informatique et systèmes embarqués</p>
              </div>
            </div>
            <div class="absolute left-1/2 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-900"></div>
            <div class="w-1/2"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Section Vidéo Voyage - Background Video -->
    <div class="relative min-h-screen flex items-center justify-center overflow-hidden">
      <!-- Vidéo en background -->
      <div class="absolute inset-0">
        <video 
          ref="videoElement"
          src="../components/Videos/Trip.mp4" 
          class="w-full h-full object-cover"
          autoplay 
          loop 
          muted
          playsinline
        ></video>
        <!-- Overlay sombre -->
        <div class="absolute inset-0 bg-gradient-to-b from-gray-900/90 via-gray-900/70 to-gray-900/90"></div>
        <div class="absolute inset-0 bg-blue-500/10"></div>
      </div>

      <!-- Effet de grille par dessus -->
      <div class="absolute inset-0 opacity-10">
        <div class="absolute inset-0 tech-pattern"></div>
      </div>

      <!-- Contenu -->
      <div class="relative z-10 max-w-5xl mx-auto px-4 py-20 text-center">
        <!-- Badge -->
        <div class="inline-block px-6 py-3 bg-blue-500/20 border border-blue-400/30 rounded-full mb-8 backdrop-blur-sm">
          <span class="text-blue-400 font-mono text-sm tracking-wider">VOYAGES & DÉCOUVERTES</span>
        </div>

        <h2 class="text-5xl md:text-6xl font-bold text-white mb-6 leading-tight">
          Explorer le <span class="bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">monde</span>
        </h2>

        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-purple-500 mx-auto mb-8"></div>

        <p class="text-xl md:text-2xl text-gray-200 leading-relaxed mb-12 max-w-3xl mx-auto">
          Milan, Italie • 2025
        </p>

        <p class="text-lg text-gray-300 leading-relaxed mb-12 max-w-3xl mx-auto">
          Entre architecture historique et modernité, chaque voyage est une source d'inspiration qui nourrit ma créativité et enrichit ma vision du monde.
        </p>


        <!-- Stats en ligne -->
        <div class="grid grid-cols-3 gap-8 max-w-2xl mx-auto">
          <div class="bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-blue-400/20 hover:border-blue-400/50 transition-all duration-300 group">
            <div class="text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent mb-2 transition-transform duration-300 group-hover:scale-110">
              15+
            </div>
            <div class="text-gray-300 text-sm uppercase tracking-wider">Pays visités</div>
          </div>
          <div class="bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-blue-400/20 hover:border-blue-400/50 transition-all duration-300 group">
            <div class="text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent mb-2 transition-transform duration-300 group-hover:scale-110">
              ∞
            </div>
            <div class="text-gray-300 text-sm uppercase tracking-wider">Découvertes</div>
          </div>
          <div class="bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-blue-400/20 hover:border-blue-400/50 transition-all duration-300 group">
            <div class="text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent mb-2 transition-transform duration-300 group-hover:scale-110">
              100%
            </div>
            <div class="text-gray-300 text-sm uppercase tracking-wider">Curiosité</div>
          </div>
        </div>
      </div>

      <!-- Coins décoratifs -->
      <div class="absolute top-8 left-8 w-20 h-20 border-t-2 border-l-2 border-blue-400/30 pointer-events-none"></div>
      <div class="absolute top-8 right-8 w-20 h-20 border-t-2 border-r-2 border-blue-400/30 pointer-events-none"></div>
      <div class="absolute bottom-8 left-8 w-20 h-20 border-b-2 border-l-2 border-blue-400/30 pointer-events-none"></div>
      <div class="absolute bottom-8 right-8 w-20 h-20 border-b-2 border-r-2 border-blue-400/30 pointer-events-none"></div>

      <!-- Particules flottantes -->
      <div class="absolute inset-0 pointer-events-none">
        <div class="floating-particle" style="top: 15%; left: 10%;"></div>
        <div class="floating-particle" style="top: 70%; right: 12%; animation-delay: 2s;"></div>
        <div class="floating-particle" style="bottom: 25%; left: 85%; animation-delay: 4s;"></div>
        <div class="floating-particle" style="top: 45%; right: 80%; animation-delay: 1s;"></div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';

export default {
  name: 'ProfilePage',
  setup() {
    const animatedYears = ref(0);
    const animatedProjects = ref(0);
    const animatedTech = ref(0);
    const animatedClients = ref(0);
    const videoElement = ref(null);

    const animateValue = (ref, target, duration = 2000) => {
      const start = 0;
      const increment = target / (duration / 16);
      let current = start;

      const timer = setInterval(() => {
        current += increment;
        if (current >= target) {
          ref.value = target;
          clearInterval(timer);
        } else {
          ref.value = Math.floor(current);
        }
      }, 16);
    };

    onMounted(() => {
      const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            animateValue(animatedYears, 5);
            animateValue(animatedProjects, 50);
            animateValue(animatedTech, 20);
            animateValue(animatedClients, 30);
            observer.disconnect();
          }
        });
      });

      const statsElement = document.querySelector('.grid.grid-cols-2');
      if (statsElement) {
        observer.observe(statsElement);
      }
    });

    return {
      animatedYears,
      animatedProjects,
      animatedTech,
      animatedClients,
      videoElement
    };
  }
};
</script>

<style scoped>
.grid-background {
  background-image: linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
                    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
  animation: gridMove 20s linear infinite;
}

@keyframes gridMove {
  0% { transform: translateY(0); }
  100% { transform: translateY(50px); }
}

/* Style minimaliste moderne */
.circle-rotate {
  transform-origin: center;
  animation: rotateCircle 20s linear infinite;
}

@keyframes rotateCircle {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

/* Petits accents discrets */
.accent-dot {
  position: absolute;
  width: 6px;
  height: 6px;
  background: linear-gradient(135deg, #3b82f6, #8b5cf6);
  border-radius: 50%;
  opacity: 0;
  transition: all 0.5s ease;
  box-shadow: 0 0 10px rgba(59, 130, 246, 0.5);
}

.group:hover .accent-dot {
  opacity: 0.6;
}

.accent-top {
  top: 0;
  left: 50%;
  transform: translateX(-50%) translateY(-10px);
}

.accent-right {
  top: 50%;
  right: 0;
  transform: translateY(-50%) translateX(10px);
}

.accent-bottom {
  bottom: 0;
  left: 50%;
  transform: translateX(-50%) translateY(10px);
}

.accent-left {
  top: 50%;
  left: 0;
  transform: translateY(-50%) translateX(-10px);
}

.tech-pattern {
  background-image: linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
                    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 30px 30px;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fadeIn 0.8s ease-out;
}

.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out 0.2s both;
}

.animate-fade-in-up-delay {
  animation: fadeInUp 0.8s ease-out 0.4s both;
}

.animate-fade-in-up-delay-1 {
  animation: fadeInUp 0.8s ease-out 0.4s both;
}

.animate-fade-in-up-delay-2 {
  animation: fadeInUp 0.8s ease-out 0.6s both;
}

.animate-pulse-slow {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.photo-card {
  transition: transform 0.3s ease;
}

.photo-card:hover {
  transform: translateY(-8px);
}

.video-showcase {
  transition: transform 0.3s ease;
}

.video-showcase:hover {
  transform: translateY(-4px);
}

.floating-particle {
  position: absolute;
  width: 6px;
  height: 6px;
  background: linear-gradient(45deg, #3b82f6, #8b5cf6);
  border-radius: 50%;
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.8);
  animation: floatParticle 6s ease-in-out infinite;
}

@keyframes floatParticle {
  0%, 100% {
    transform: translateY(0) translateX(0);
    opacity: 0.3;
  }
  25% {
    transform: translateY(-30px) translateX(15px);
    opacity: 0.8;
  }
  50% {
    transform: translateY(-60px) translateX(-10px);
    opacity: 0.5;
  }
  75% {
    transform: translateY(-30px) translateX(-20px);
    opacity: 0.8;
  }
}
</style>