resource "aws_ecr_repository" "ecr" {
    name = "${var.teamname}-repository"
    image_tag_mutability = var.image_mutability
    image_scanning_configuration {
        scan_on_push = true
    }
    encryption_configuration {
        encryption_type = "AES256"
    }
}