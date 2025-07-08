function handleReister() {
    // 1、收集表单数据
    const username = document.getElementById('regUsername').value.trim();
    const password = document.getElementById('regPassword').value.trim();
    const confirmPassword = document.getElementById('confirmPassword').value.trim();

    const button = document.getElementById('registerForm').querySelector('button');
    const buttonText = document.getElementById('buttonText');
    const spinner = document.getElementById('spinner');

    // 显示加载状态
    button.disabled = true;
    buttonText.style.display = 'none';
    spinner.style.display = 'block';

    // 前端基础验证
    // 不为空、两次输入的密码不匹配
    let hasError = false;
    if (username === '') {
        document.getElementById('usernameError').style.display = 'block';
        document.getElementById('regUsername').style.borderColor = '#ff6b81';
        hasError = true;
    } else {
        document.getElementById('usernameError').style.display = 'none';
        document.getElementById('regUsername').style.borderColor = 'rgba(0, 0, 0, 0.1)';
    }

    if (password === '') {
        document.getElementById('passwordError').style.display = 'block';
        document.getElementById('regPassword').style.borderColor = '#ff6b81';
        hasError = true;
    } else {
        document.getElementById('passwordError').style.display = 'none';
        document.getElementById('regPassword').style.borderColor = 'rgba(0, 0, 0, 0.1)';
    }

    if (confirmPassword === '') {
        document.getElementById('confirmPasswordError').style.display = 'block';
        document.getElementById('confirmPassword').style.borderColor = '#ff6b81';
        hasError = true;
    } else if (confirmPassword !== password) {
        document.getElementById('confirmPasswordError').textContent = '两次输入的密码不一致';
        document.getElementById('confirmPasswordError').style.display = 'block';
        document.getElementById('confirmPassword').style.borderColor = '#ff6b81';
        hasError = true;
    } else {
        document.getElementById('confirmPasswordError').style.display = 'none';
        document.getElementById('confirmPassword').style.borderColor = 'rgba(0, 0, 0, 0.1)';
    }

    if (hasError) {
        // 恢复按钮状态（仅在验证失败时）
        button.disabled = false;
        buttonText.style.display = 'inline';
        spinner.style.display = 'none';
        return; // 终止函数，不发送请求
    }

    // 发送真实 AJAX 请求
    fetch('/api/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            // 'X-CSRF-Token': '{{ csrf_token }}', // 启用 CSRF 保护
        },
        body: JSON.stringify({
            username: username,
            password: password,
            confirmPassword: confirmPassword
        })
    })
        .then(response => response.json())
        .then(data => {
            // 处理后端返回结果
            if (data.code === 0) {
                // 注册成功
                alert('注册成功！即将跳转');
                window.location.href = '/dashboard_login';
            } else {
                // 注册失败，显示后端错误
                showBackendError(data.msg);
            }
        })
        .catch(error => {
            console.error('网络错误:', error);
            showBackendError('网络请求失败，请重试');
        })
        .finally(() => {
            // 恢复按钮状态（无论成功失败）
            button.disabled = false;
            buttonText.style.display = 'inline';
            spinner.style.display = 'none';
        });
}

function showBackendError(msg) {
    alert(msg);
}