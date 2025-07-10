document.addEventListener('DOMContentLoaded', function () {
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
            "line_linked": { "enable": true, "opacity": 0.1 },
            "z-index": { "value": 1 }
        }
    });
}

// 按钮波纹
function bindBtnRipple() {
    const buttons = document.querySelectorAll('.btn-primary, .btn-warning, .btn-danger');
    buttons.forEach(btn => {
        // 提前创建波纹元素
        const ripple = document.createElement('div');
        ripple.className = 'btn-ripple';
        btn.appendChild(ripple);

        btn.addEventListener('click', function (e) {
            const rect = btn.getBoundingClientRect();
            const x = e.clientX - rect.left;
            const y = e.clientY - rect.top;

            ripple.style.left = x + 'px';
            ripple.style.top = y + 'px';
            ripple.style.width = '0';
            ripple.style.height = '0';
            ripple.style.opacity = '1'; // 确保可见

            setTimeout(() => {
                const size = Math.max(rect.width, rect.height) * 2;
                ripple.style.width = size + 'px';
                ripple.style.height = size + 'px';
                ripple.style.opacity = '0'; // 淡出效果
            }, 10);
        });
    });
}

// 加载用户列表
function loadUsers() {
    fetch('/api/users', {
        method: 'PUT',
        credentials: 'include',
    })
        .then(response => {
            // alert("1111")
            if (!response.ok) {
                throw response.json();
            }
            // alert("22")
            return response.json();
        })
        .then(data => {
                // alert("3333");
            renderUsers(data.users);
            // alert("444")
        }
        )
        .catch(error => {
            console.error("加载用户列表失败", error);
            alert("加载用户列表失败，请刷新页面重试");
        });
}

// 渲染用户列表
function renderUsers(users) {
    const tbody = document.querySelector('tbody');
    tbody.innerHTML = '';
    users.forEach(user => {
        const tr = document.createElement('tr');
        tr.className = 'table-row-hover';

        // 转换日期格式
        const regDate = new Date(user.created_at).toLocaleDateString('zh-CN', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit'
        });

        tr.innerHTML = `
            <td>${user.id}</td> 
            <td>${user.username}</td>
            <td>
                ${user.role === '0' ?
            '<span class="badge badge-admin">管理员</span>' :
            '<span class="badge badge-user">普通用户</span>'}
            </td>
            <td>${regDate}</td>
            ${document.querySelector('.admin-badge') ? `
            <td>
                <div class="btn-group">
                    <button class="btn-warning" onclick="openEditModal(${user.id}, '${user.username}', '${user.role}')">
                        <i class="bi bi-pencil me-1"></i> 编辑
                        <div class="btn-ripple"></div>
                    </button>
                    <button class="btn-danger" onclick="deleteUser(${user.id})">
                        <i class="bi bi-trash me-1"></i> 删除
                        <div class="btn-ripple"></div>
                    </button>
                </div>
            </td>
            ` : ''}
        `;
        tbody.appendChild(tr);
    })
}
// 删除用户
function deleteUser(userId) {
    if (!confirm(`确定要删除ID为${userId}的用户吗`)) {
        return;
    }
    fetch(`/api/users/${userId}`, {
        method: 'DELETE'
    })
        .then(response => {
            if (response.ok) {
                loadUsers(); // 刷新用户列表
                alert('用户删除成功');
            } else {
                alert('删除用户失败');
            }
        })
        .catch(error => {
            console.error('删除用户失败', error);
            alert('删除用户失败');
        })
}

// 管理员特殊权限

// 打开编辑模态框
function openEditModal(id, username, role) {
    document.getElementById('editUserId').value = id;
    document.getElementById('editUsername').value = username;
    document.getElementById('editIsAdmin').checked = role === 'admin';

    // 使用Bootstrap显示模态框
    const editModal = new bootstrap.Modal(document.getElementById('editUserModal'));
    editModal.show();
}

// 提交编辑表单
function submitEditUser() {
    const userId = document.getElementById('editUserId').value;
    const username = document.getElementById('editUsername').value;
    const isAdmin = document.getElementById('editIsAdmin').checked;
    const role = isAdmin ? '0' : '1';

    fetch(`/api/users/${userId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            role: role
        })
    })
        .then(response => {
            alert("正在尝试更改")
            if (response.ok) {
                loadUsers(); // 刷新用户列表
                // 隐藏模态框
                const editModal = bootstrap.Modal.getInstance(document.getElementById('editUserModal'));
                editModal.hide();
                alert('用户更新成功');
            } else {
                alert('更新用户失败');
            }
        })
        .catch(error => {
            console.error('更新用户失败:', error);
            alert('更新用户失败');
        });

    return false; // 阻止表单默认提交
}

// 提交添加用户表单（需要后端接口支持）
function submitAddUser() {
    const username = document.getElementById('addUsername').value;
    const password = document.getElementById('addPassword').value;
    const isAdmin = document.getElementById('addIsAdmin').checked;
    const role = isAdmin ? '0' : '1';
    fetch('/api/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            password: password,
            role: role
        })
    })
        .then(response => {
            if (response.ok) {
                loadUsers(); // 刷新用户列表
                // 隐藏模态框
                const addModal = bootstrap.Modal.getInstance(document.getElementById('addUserModal'));
                addModal.hide();
                alert('用户添加成功');
            } else {
                alert('添加用户失败');
            }
        })
        .catch(error => {
            console.error('添加用户失败:', error);
            alert('添加用户失败');
        });

    return false; // 阻止表单默认提交
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