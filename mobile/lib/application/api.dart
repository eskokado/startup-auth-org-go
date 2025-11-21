import 'dart:convert';
import 'package:http/http.dart' as http;
import '../domain/entities.dart';

class AuthResult {
  final String accessToken;
  final Map<String, dynamic> user;
  AuthResult(this.accessToken, this.user);
}

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

  Future<AuthResult> login(String email, String password) async {
    final res = await http.post(Uri.parse('$baseUrl/auth/login'), headers: {'Content-Type': 'application/json'}, body: jsonEncode({'email': email, 'password': password}));
    final data = jsonDecode(res.body);
    return AuthResult(data['access_token'], data['user'] as Map<String, dynamic>);
  }

  Future<void> register(String name, String email, String password, String imageUrl) async {
    await http.post(Uri.parse('$baseUrl/auth/register'), headers: {'Content-Type': 'application/json'}, body: jsonEncode({'name': name, 'email': email, 'password': password, 'password_confirmation': password, 'image_url': imageUrl}));
  }

  Future<void> forgotPassword(String email, String redirectUrl) async {
    await http.post(Uri.parse('$baseUrl/auth/forgot-password'), headers: {'Content-Type': 'application/json'}, body: jsonEncode({'email': email, 'redirect_url': redirectUrl}));
  }

  Future<void> updateName(String userId, String name) async {
    await http.put(Uri.parse('$baseUrl/user/name/$userId'), headers: _headers, body: jsonEncode({'name': name}));
  }

  Future<void> updatePassword(String userId, String password) async {
    await http.put(Uri.parse('$baseUrl/user/password/$userId'), headers: _headers, body: jsonEncode({'password': password}));
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