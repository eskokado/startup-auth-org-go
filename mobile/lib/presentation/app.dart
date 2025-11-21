import 'package:flutter/material.dart';
import '../application/api.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import '../domain/entities.dart';
import 'login.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'profile.dart';

class AppRoot extends StatefulWidget {
  final String baseUrl;
  const AppRoot({super.key, required this.baseUrl});
  @override
  State<AppRoot> createState() => _AppRootState();
}

class _AppRootState extends State<AppRoot> {
  Organization? org;
  List<TaskEntity> tasks = [];
  final titleCtrl = TextEditingController();
  final descCtrl = TextEditingController();
  String? token;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    final prefs = await SharedPreferences.getInstance();
    final tk = prefs.getString('access-token');
    final exp = int.tryParse(prefs.getString('token-exp') ?? '') ?? 0;
    final now = DateTime.now().millisecondsSinceEpoch ~/ 1000;
    if (tk == null || exp <= now) {
      await _login();
      return;
    }
    final userId = prefs.getString('user-id') ?? '';
    token = tk;
    final api = ApiClient(widget.baseUrl, token: token);
    final o = await api.getPersonalOrg(userId);
    final t = await api.listTasks(o.id);
    setState(() { org = o; tasks = t; });
  }

  Future<void> _createTask() async {
    if (org == null || titleCtrl.text.isEmpty) return;
    final api = ApiClient(widget.baseUrl, token: token);
    await api.createTask(org!.id, titleCtrl.text, descCtrl.text);
    titleCtrl.clear();
    descCtrl.clear();
    await _load();
  }

  Future<void> _login() async {
    final tk = await Navigator.push(context, MaterialPageRoute(builder: (_) => LoginPage(baseUrl: widget.baseUrl)));
    if (tk is String) { setState(() { token = tk; }); await _load(); }
  }

  Future<void> _logout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.clear();
    setState(() { token = null; org = null; tasks = []; });
    await _login();
  }

  Future<void> _checkout(String cycle) async {
    if (org == null) return;
    final url = Uri.parse('${widget.baseUrl}/billing/checkout');
    final res = await http.post(url, headers: {'Content-Type': 'application/json', if (token != null) 'Authorization': 'Bearer $token'}, body: jsonEncode({'organization_id': org!.id, 'plan': 'PERSONAL', 'cycle': cycle, 'success_url': 'app://success', 'cancel_url': 'app://cancel'}));
    final data = jsonDecode(res.body);
    final u = Uri.parse(data['checkout_url']);
    if (await canLaunchUrl(u)) { await launchUrl(u, mode: LaunchMode.externalApplication); }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(title: const Text('Tasks'), actions: [
          IconButton(onPressed: _login, icon: const Icon(Icons.login)),
          IconButton(onPressed: () => Navigator.push(context, MaterialPageRoute(builder: (_) => ProfilePage(baseUrl: widget.baseUrl))), icon: const Icon(Icons.person)),
          IconButton(onPressed: _logout, icon: const Icon(Icons.logout)),
        ]),
        body: Padding(
          padding: const EdgeInsets.all(16),
          child: Column(children: [
            Text('Org: ${org?.name ?? '...'}'),
            Row(children: [
              ElevatedButton(onPressed: () => _checkout('MONTHLY'), child: const Text('Assinar Mensal')),
              const SizedBox(width: 8),
              ElevatedButton(onPressed: () => _checkout('SEMIANNUAL'), child: const Text('Assinar Semestral')),
              const SizedBox(width: 8),
              ElevatedButton(onPressed: () => _checkout('ANNUAL'), child: const Text('Assinar Anual')),
            ]),
            Row(children: [
              Expanded(child: TextField(controller: titleCtrl, decoration: const InputDecoration(labelText: 'Title'))),
              const SizedBox(width: 8),
              Expanded(child: TextField(controller: descCtrl, decoration: const InputDecoration(labelText: 'Description'))),
              const SizedBox(width: 8),
              ElevatedButton(onPressed: _createTask, child: const Text('Create')),
            ]),
            const SizedBox(height: 16),
            Expanded(
              child: ListView.builder(
                itemCount: tasks.length,
                itemBuilder: (_, i) {
                  final t = tasks[i];
                  return ListTile(title: Text(t.title), subtitle: Text(t.description), trailing: Text(t.status));
                },
              ),
            )
          ]),
        ),
      ),
    );
  }
}