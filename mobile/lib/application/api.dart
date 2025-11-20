import 'dart:convert';
import 'package:http/http.dart' as http;
import '../domain/entities.dart';

class ApiClient {
  final String baseUrl;
  final String? token;
  ApiClient(this.baseUrl, {this.token});

  Map<String, String> get _headers => {
        'Content-Type': 'application/json',
        if (token != null) 'Authorization': 'Bearer $token',
      };

  Future<Organization> getPersonalOrg(String ownerId) async {
    final res = await http.get(Uri.parse('$baseUrl/org/personal/$ownerId'), headers: _headers);
    final j = jsonDecode(res.body);
    return Organization.fromJson(j);
  }

  Future<List<TaskEntity>> listTasks(String orgId) async {
    final res = await http.get(Uri.parse('$baseUrl/tasks?organization_id=$orgId'), headers: _headers);
    final list = jsonDecode(res.body) as List;
    return list.map((e) => TaskEntity.fromJson(e)).toList();
  }

  Future<void> createTask(String orgId, String title, String description) async {
    await http.post(Uri.parse('$baseUrl/tasks'), headers: _headers, body: jsonEncode({'organization_id': orgId, 'title': title, 'description': description}));
  }
}