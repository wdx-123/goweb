// 三个交互效果
document.addEventListener('DOMContentLoaded', function() {
    initParticles();  // 粒子背景动画
    initTypewriter(); // 卡片描述
    bindBtnRipple();  // 按钮点击波纹效果
});

// 粒子背景（颜色更柔和）
function initParticles() {
    particlesJS('particles-js', {
        "particles": {      
            "number": { "value": 50, "density": { "enable": true, "value_area": 800 } },
            "color": { "value": '#4299e1' }, // 柔和蓝色粒子
            "shape": { "type": "circle" },
            "opacity": {
                "value": 0.3, // 降低透明度避免干扰
                "random": true
            },
            "size": { "value": 3, "random": true },
            "line_linked": {
                "enable": true,
                "distance": 150,
                "color": "#4299e1",
                "opacity": 0.2,
                "width": 1
            },
            "move": { "enable": true, "speed": 1 }
        },
        "interactivity": {
            "onhover": { "enable": true, "mode": "grab" },
            "onclick": { "enable": true, "mode": "push" }
        }
    });
}

// 打字机效果（速度适中）
function initTypewriter() {
    const descElements = document.querySelectorAll('.card-desc');
    descElements.forEach(el => {
        const text = el.getAttribute('data-text');
        let index = 0;
        el.textContent = '';

        const typeInterval = setInterval(() => {
            if (index < text.length) {
                el.textContent += text.charAt(index);
                index++;
            } else {
                clearInterval(typeInterval);
            }
        }, 80); // 适中速度，提升可读性
    });
}

// 按钮波纹
function bindBtnRipple() {
    const buttons = document.querySelectorAll('.card-btn');
    buttons.forEach(btn => {
        btn.addEventListener('click', function(e) {
            const rect = btn.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;

            const ripple = this.querySelector('.btn-ripple');
            ripple.style.left = x + 'px';
            ripple.style.top = y + 'px';
            ripple.style.width = '0';
            ripple.style.height = '0';

            setTimeout(() => {
                ripple.style.width = Math.max(rect.width, rect.height) * 2 + 'px';
                ripple.style.height = Math.max(rect.width, rect.height) * 2 + 'px';
            }, 10);
        });
    });
}


async function logoutAndRedirect() {
    try {
        // 显示加载状态（可选）
        const logoutBtn = document.querySelector('.logout-btn');
        if (logoutBtn) {
            logoutBtn.innerHTML = '<i class="bi bi-circle-fill text-danger me-1"></i>退出中...';
            logoutBtn.disabled = true;
        }
        // 调用退出登录 API
        const response = await fetch('/api/sessions', {
            method: 'DELETE', // 假设退出登录使用 DELETE 方法
            credentials: 'include', // 包含 cookie 等凭证
        });
        if (!response.ok) {
            throw new Error(`退出登录失败: ${response.statusText}`);
        }
        // 退出成功后重定向到登录页或首页
        window.location.href = '/login'; // 修改为你想要跳转的路径
    } catch (error) {
        console.error('退出登录过程中出错:', error);
        alert('退出登录失败，请稍后再试');

        // 恢复按钮状态
        const logoutBtn = document.querySelector('.logout-btn');
        if (logoutBtn) {
            logoutBtn.innerHTML = '<i class="bi bi-power text-danger me-1"></i>退出';
            logoutBtn.disabled = false;
        }
    }
}