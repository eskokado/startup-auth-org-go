import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

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
    final res = await http.post(Uri.parse('${widget.baseUrl}/auth/login'), headers: {'Content-Type': 'application/json'}, body: jsonEncode({'email': emailCtrl.text, 'password': passCtrl.text}));
    final data = jsonDecode(res.body);
    setState(() { token = data['access_token']; });
    if (context.mounted) Navigator.pop(context, token);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Login')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(children: [
          TextField(controller: emailCtrl, decoration: const InputDecoration(labelText: 'Email')),
          TextField(controller: passCtrl, decoration: const InputDecoration(labelText: 'Password'), obscureText: true),
          const SizedBox(height: 12),
          ElevatedButton(onPressed: _login, child: const Text('Entrar')),
        ]),
      ),
    );
  }
}