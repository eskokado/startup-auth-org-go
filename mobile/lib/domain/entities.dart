class Organization {
  final String id;
  final String name;
  Organization({required this.id, required this.name});
  factory Organization.fromJson(Map<String, dynamic> j) => Organization(id: j['id'], name: j['name']);
}

class TaskEntity {
  final String id;
  final String title;
  final String description;
  final String status;
  TaskEntity({required this.id, required this.title, required this.description, required this.status});
  factory TaskEntity.fromJson(Map<String, dynamic> j) => TaskEntity(id: j['id'], title: j['title'], description: j['description'] ?? '', status: j['status']);
}