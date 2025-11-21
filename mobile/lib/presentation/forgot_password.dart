import 'package:flutter/material.dart';
import '../application/api.dart';

class ForgotPasswordPage extends StatefulWidget {
  final String baseUrl;
  const ForgotPasswordPage({super.key, required this.baseUrl});
  @override
  State<ForgotPasswordPage> createState() => _ForgotPasswordPageState();
}

class _ForgotPasswordPageState extends State<ForgotPasswordPage> {
  final emailCtrl = TextEditingController();

  Future<void> _send() async {
    final api = ApiClient(widget.baseUrl);
    await api.forgotPassword(emailCtrl.text, 'app://reset');
    if (context.mounted) Navigator.pop(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Esqueci a senha')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(children: [
          TextField(controller: emailCtrl, decoration: const InputDecoration(labelText: 'Email')),
          const SizedBox(height: 12),
          ElevatedButton(onPressed: _send, child: const Text('Enviar link')),
        ]),
      ),
    );
  }
}