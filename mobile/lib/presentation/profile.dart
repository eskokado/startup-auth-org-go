import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../application/api.dart';

class ProfilePage extends StatefulWidget {
  final String baseUrl;
  const ProfilePage({super.key, required this.baseUrl});
  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  final nameCtrl = TextEditingController();
  final passCtrl = TextEditingController();
  String userId = '';
  String? token;

  @override
  void initState() {
    super.initState();
    _init();
  }

  Future<void> _init() async {
    final prefs = await SharedPreferences.getInstance();
    userId = prefs.getString('user-id') ?? '';
    nameCtrl.text = prefs.getString('user-name') ?? '';
    token = prefs.getString('access-token');
    setState(() {});
  }

  Future<void> _updateName() async {
    if (userId.isEmpty || nameCtrl.text.isEmpty) return;
    final api = ApiClient(widget.baseUrl, token: token);
    await api.updateName(userId, nameCtrl.text);
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('user-name', nameCtrl.text);
  }

  Future<void> _updatePassword() async {
    if (userId.isEmpty || passCtrl.text.isEmpty) return;
    final api = ApiClient(widget.baseUrl, token: token);
    await api.updatePassword(userId, passCtrl.text);
    passCtrl.clear();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Meu Perfil')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(children: [
          TextField(controller: nameCtrl, decoration: const InputDecoration(labelText: 'Nome')),
          const SizedBox(height: 8),
          ElevatedButton(onPressed: _updateName, child: const Text('Salvar nome')),
          const Divider(height: 32),
          TextField(controller: passCtrl, decoration: const InputDecoration(labelText: 'Nova senha'), obscureText: true),
          const SizedBox(height: 8),
          ElevatedButton(onPressed: _updatePassword, child: const Text('Alterar senha')),
        ]),
      ),
    );
  }
}