import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../application/api.dart';
import 'register.dart';
import 'forgot_password.dart';

class LoginPage extends StatefulWidget {
  final String baseUrl;
  const LoginPage({super.key, required this.baseUrl});
  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final emailCtrl = TextEditingController();
  final passCtrl = TextEditingController();
  String? token;

  Future<void> _login() async {
    final api = ApiClient(widget.baseUrl);
    final result = await api.login(emailCtrl.text, passCtrl.text);
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('access-token', result.accessToken);
    await prefs.setString('user-id', (result.user['id'] ?? '') as String);
    await prefs.setString('user-name', (result.user['name'] ?? '') as String);
    await prefs.setString('user-email', (result.user['email'] ?? '') as String);
    // Decode JWT for org_id and exp
    try {
      final parts = result.accessToken.split('.');
      if (parts.length == 3) {
        final decoded = jsonDecode(utf8.decode(base64Url.decode(base64Url.normalize(parts[1]))));
        final orgId = decoded['org_id']?.toString() ?? '';
        final exp = decoded['exp']?.toString() ?? '';
        final plan = decoded['plan']?.toString() ?? '';
        if (orgId.isNotEmpty) await prefs.setString('org-id', orgId);
        if (exp.isNotEmpty) await prefs.setString('token-exp', exp);
        if (plan.isNotEmpty) await prefs.setString('plan', plan);
      }
    } catch (_) {}
    setState(() { token = result.accessToken; });
    if (context.mounted) Navigator.pop(context, token);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(24),
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            const Text('Login', style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
            const SizedBox(height: 16),
            TextField(controller: emailCtrl, decoration: const InputDecoration(labelText: 'Email')),
            const SizedBox(height: 8),
            TextField(controller: passCtrl, decoration: const InputDecoration(labelText: 'Senha'), obscureText: true),
            const SizedBox(height: 16),
            ElevatedButton(onPressed: _login, child: const Text('Entrar')),
            const SizedBox(height: 12),
            Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
              TextButton(onPressed: () { Navigator.push(context, MaterialPageRoute(builder: (_) => RegisterPage(baseUrl: widget.baseUrl))); }, child: const Text('Registrar')),
              TextButton(onPressed: () { Navigator.push(context, MaterialPageRoute(builder: (_) => ForgotPasswordPage(baseUrl: widget.baseUrl))); }, child: const Text('Esqueci a senha')),
            ])
          ]),
        ),
      ),
    );
  }
}