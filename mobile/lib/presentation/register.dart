import 'package:flutter/material.dart';
import '../application/api.dart';

class RegisterPage extends StatefulWidget {
  final String baseUrl;
  const RegisterPage({super.key, required this.baseUrl});
  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final nameCtrl = TextEditingController();
  final emailCtrl = TextEditingController();
  final passCtrl = TextEditingController();

  Future<void> _register() async {
    final api = ApiClient(widget.baseUrl);
    await api.register(nameCtrl.text, emailCtrl.text, passCtrl.text, '');
    if (context.mounted) Navigator.pop(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Registrar')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(children: [
          TextField(controller: nameCtrl, decoration: const InputDecoration(labelText: 'Nome')),
          TextField(controller: emailCtrl, decoration: const InputDecoration(labelText: 'Email')),
          TextField(controller: passCtrl, decoration: const InputDecoration(labelText: 'Senha'), obscureText: true),
          const SizedBox(height: 12),
          ElevatedButton(onPressed: _register, child: const Text('Criar conta')),
        ]),
      ),
    );
  }
}