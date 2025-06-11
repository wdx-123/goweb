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

document.getElementById('loginForm').addEventListener('submit',  async function (e) {
    e.preventDefault(); // 阻止表单默认刷新行为
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const rememberMe = document.getElementById('rememberMe').checked; // 记住界面
    const button = this.querySelector('button');
    const buttonText = document.getElementById('buttonText');
    const spinner = document.getElementById('spinner'); // 登入界面
    const globalError = document.getElementById('globalError');

    // 重置状态：隐藏全局错误 + 显示加载
    globalError.style.display = 'none';
    button.disabled = true;
    buttonText.style.display = 'none';
    spinner.style.display = 'inline-block';


    // 基础表单验证（空值检查）
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
        // 发送ajax请求到后端 /login接口
        try {
           // todo
            const response = await fetch('/api/login',{ // 请求页面
                method:'POST', // 登入采用POST更加安全
                headers:{
                    'Content-Type':'application/json', // 告诉后端数据格式是json
                },
                body:JSON.stringify({
                    username:username,
                    password:password,
                    rememberMe:rememberMe // boolean格式
                }),
                credentials:'include' // 支持 Cookie 跨域
                }
            );

            // 解析 后端返回的JSON
            const result = await response.json();
            // 回复按钮状态
            button.disabled = false;
            buttonText.style.display = 'inline';
            spinner.style.display = 'none';

            // 处理响应结果
            if(response.ok && result.code === 0){
                // 登入成功，后端已经设置Cookie，跳转页面
                window.location.href = '/dashboard_login';
            }else{ // 登入失败
                globalError.textContent  = result.error || '登入失败，请重试';
                globalError.style.display = 'block'; // 显示错误

            }
        } catch (err) {
            // 网络错误（如后端没启动）
            button.disabled = false;
            buttonText.style.display = 'inline';
            spinner.style.display = 'none';
            globalError.textContent = '网络错误，请检查服务器';
            globalError.style.display = 'block';
        }

    } else {
        // 恢复按钮状态
        button.disabled = false;
        buttonText.style.display = 'inline';
        spinner.style.display = 'none';
    }
});