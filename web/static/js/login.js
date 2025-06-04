// 密码显示/隐藏功能
function togglePasswordVisibility() {
    const passwordInput = document.getElementById('password');
    const eyeIcon = document.getElementById('eyeIcon');
    if (passwordInput.type === 'password') {
        passwordInput.type = 'text';
        eyeIcon.classList.remove('fa-eye');
        eyeIcon.classList.add('fa-eye-slash');
    } else {
        passwordInput.type = 'password';
        eyeIcon.classList.remove('fa-eye-slash');
        eyeIcon.classList.add('fa-eye');
    }
}

// 表单提交处理
document.getElementById('loginForm').addEventListener('submit', function (e) {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const button = this.querySelector('button');
    const buttonText = document.getElementById('buttonText');
    const spinner = document.getElementById('spinner');

    // 显示加载状态
    button.disabled = true;
    buttonText.style.display = 'none';
    spinner.style.display = 'block';

    // 简单的表单验证
    let hasError = false;
    if (username === '') {
        document.getElementById('usernameError').style.display = 'block';
        document.getElementById('username').style.borderColor = '#ff6b81';
        hasError = true;
    } else {
        document.getElementById('usernameError').style.display = 'none';
        document.getElementById('username').style.borderColor = 'rgba(0, 0, 0, 0.1)';
    }

    if (password === '') {
        document.getElementById('passwordError').style.display = 'block';
        document.getElementById('password').style.borderColor = '#ff6b81';
        hasError = true;
    } else {
        document.getElementById('passwordError').style.display = 'none';
        document.getElementById('password').style.borderColor = 'rgba(0, 0, 0, 0.1)';
    }

    if (!hasError) {
        // 模拟异步请求（实际需替换为真实接口）
        setTimeout(() => {
            button.disabled = false;
            buttonText.style.display = 'inline';
            spinner.style.display = 'none';
            // 跳转页面
            window.location.href = '/dashboard';
        }, 2000);
    } else {
        // 恢复按钮状态
        button.disabled = false;
        buttonText.style.display = 'inline';
        spinner.style.display = 'none';
    }
});