document.addEventListener('DOMContentLoaded', function() {
    initParticles();
    bindBtnRipple();
});

// 粒子背景（柔和模式）
function initParticles() {
    particlesJS('particles-js', {
        "particles": {
            "number": { "value": 40, "density": { "enable": true, "value_area": 800 } },
            "color": { "value": "#4299e1" },
            "opacity": { "value": 0.2 },
            "size": { "value": 2 },
            "line_linked": { "enable": true, "opacity": 0.1 }
        }
    });
}

// 按钮波纹
function bindBtnRipple() {
    const buttons = document.querySelectorAll('.btn-primary, .btn-warning, .btn-danger');
    buttons.forEach(btn => {
        btn.addEventListener('click', function(e) {
            const rect = btn.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;

            let ripple = this.querySelector('.btn-ripple');
            if (!ripple) {
                ripple = document.createElement('div');
                ripple.className = 'btn-ripple';
                this.appendChild(ripple);
            }

            ripple.style.left = x + 'px';
            ripple.style.top = y + 'px';
            ripple.style.width = '0';
            ripple.style.height = '0';

            setTimeout(() => {
                const size = Math.max(rect.width, rect.height) * 2;
                ripple.style.width = size + 'px';
                ripple.style.height = size + 'px';
            }, 10);
        });
    });
}